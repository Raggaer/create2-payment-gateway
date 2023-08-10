package main

import (
	"context"
	"create2-payment-gateway/controllers"
	"create2-payment-gateway/deployer"
	"create2-payment-gateway/models"
	"crypto/ecdsa"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

var (
	deployerNonce uint64
)

// RunScheduler runs the scheduler checking for payments address balances
func RunScheduler(interval time.Duration, client *ethclient.Client, pk *ecdsa.PrivateKey, db *sql.DB) {
	deployer, err := deployer.NewContractFactory(common.HexToAddress(viper.GetString("factory_addr")), client)
	if err != nil {
		panic(err)
	}

	// Setup initial nonce
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(pk.PublicKey))
	if err != nil {
		panic(err)
	}
	atomic.StoreUint64(&deployerNonce, nonce)

	fmt.Println("Starting scheduler with nonce", deployerNonce)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if err := schedulerTick(deployer, client, pk, db); err != nil {
			fmt.Println("Error in scheduler tick:", err)
			continue
		}
	}
}

func schedulerTick(deployer *deployer.ContractFactory, client *ethclient.Client, pk *ecdsa.PrivateKey, db *sql.DB) error {
	payments, err := models.RetrieveWaitingPayments(db)
	if err != nil {
		return fmt.Errorf("error retrieving waiting payments: %v", err)
	}
	if len(payments) == 0 {
		return nil
	}

	// Check for each payment
	for _, payment := range payments {
		shouldDeploy, err := checkPayment(client, payment)
		if err != nil {
			return fmt.Errorf("error checking payment: %v", err)
		}

		if !shouldDeploy {
			continue
		}
		if err := deployDrain(deployer, client, pk, payment); err != nil {
			return fmt.Errorf("error deploying drain: %v", err)
		}
	}
	return nil
}

func deployDrain(deployer *deployer.ContractFactory, client *ethclient.Client, pk *ecdsa.PrivateKey, payment models.Payment) error {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(1337))
	if err != nil {
		return fmt.Errorf("error creating transactor: %v", err)
	}
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	auth.GasLimit = 1000000
	auth.Nonce = big.NewInt(int64(atomic.LoadUint64(&deployerNonce)))

	// Convert salt
	salt, err := hexSaltToByte32(payment.Salt)
	if err != nil {
		return fmt.Errorf("error decoding salt: %v", err)
	}

	// Init hash bytecode
	_, bc, err := controllers.GenerateInitHash(viper.GetString("owner_addr"))
	if err != nil {
		return fmt.Errorf("error generating init hash: %v", err)
	}

	// Deploy drain
	tx, err := deployer.Deploy(auth, bc, salt)
	if err != nil {
		return fmt.Errorf("error deploying drain: %v", err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}
	atomic.AddUint64(&deployerNonce, 1)

	deployedTo, err := checkDeployedDrain(receipt)
	if err != nil {
		return fmt.Errorf("error checking deployed drain: %v", err)
	}
	if deployedTo != payment.Address {
		fmt.Printf("WARNING: Deployed drain to different address than expected %s != %s\n", deployedTo, payment.Address)
	}
	return nil
}

func checkDeployedDrain(receipt *types.Receipt) (string, error) {
	if len(receipt.Logs) == 0 {
		return "", errors.New("no logs in receipt")
	}

	// First log is the deployed drain address
	drainAddr := common.BytesToAddress(receipt.Logs[0].Data)
	return drainAddr.Hex(), nil
}

func checkPayment(client *ethclient.Client, payment models.Payment) (bool, error) {
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(payment.Address), nil)
	if err != nil {
		return false, fmt.Errorf("error retrieving balance: %v", err)
	}

	paymentAmount := big.NewInt(int64(payment.Amount))
	paymentAmount.Mul(paymentAmount, big.NewInt(1e18))

	return balance.Cmp(paymentAmount) >= 0, nil
}

func parsePrivateKey(pk string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(pk)
}

func hexSaltToByte32(salt string) ([32]byte, error) {
	s, err := hex.DecodeString(salt)
	if err != nil {
		return [32]byte{}, err
	}
	if len(s) != 32 {
		return [32]byte{}, fmt.Errorf("salt must be 32 bytes")
	}

	var b [32]byte
	copy(b[:], s)
	return b, nil
}
