package controllers

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"create2-payment-gateway/models"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

var (
	drainABI, _ = abi.JSON(strings.NewReader(`
  [
    {
      "inputs": [
        {
          "internalType": "address payable",
          "name": "_owner",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "stateMutability": "payable",
      "type": "receive"
    }
  ]
  `))

	drainBytecode = common.FromHex("0x608060405234801561001057600080fd5b5060405161035b38038061035b833981810160405281019061003291906101b3565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600047111561014a5760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516100c290610211565b60006040518083038185875af1925050503d80600081146100ff576040519150601f19603f3d011682016040523d82523d6000602084013e610104565b606091505b5050905080610148576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013f90610283565b60405180910390fd5b505b506102a3565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061018082610155565b9050919050565b61019081610175565b811461019b57600080fd5b50565b6000815190506101ad81610187565b92915050565b6000602082840312156101c9576101c8610150565b5b60006101d78482850161019e565b91505092915050565b600081905092915050565b50565b60006101fb6000836101e0565b9150610206826101eb565b600082019050919050565b600061021c826101ee565b9150819050919050565b600082825260208201905092915050565b7f447261696e3a204661696c656420746f2073656e642045746865720000000000600082015250565b600061026d601b83610226565b915061027882610237565b602082019050919050565b6000602082019050818103600083015261029c81610260565b9050919050565b60aa806102b16000396000f3fe608060405236606f5760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015606d573d6000803e3d6000fd5b005b600080fdfea2646970667358221220a1180c773644b04aaa2d7d31b9f581741d399ced2d44a63c934c66824448248a64736f6c63430008120033")
)

type createPaymentResponse struct {
	Address string `json:"address"`
}

type createPaymentRequest struct {
	Amount int `json:"amount"`
}

// CreatePayment is a controller that handles the payment creation
func CreatePayment(w http.ResponseWriter, req *http.Request, db *sql.DB) error {
	// Derive the contract address
	addr, salt, err := derivePaymentAddress(
		viper.GetString("factory_addr"),
		viper.GetString("owner_addr"),
		"",
	)
	if err != nil {
		return fmt.Errorf("could not derive payment address: %v", err)
	}

	v := addr.Hex()

	// Create the payment
	if _, err := models.CreatePayment(salt[2:], 20, v, db); err != nil {
		return fmt.Errorf("could not create payment: %v", err)
	}

	response := createPaymentResponse{
		Address: v,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func readPaymentRequest(w http.ResponseWriter, req *http.Request) (*createPaymentRequest, error) {
	req.Body = http.MaxBytesReader(w, req.Body, 1024)

	var request createPaymentRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("could not read request: %v", err)
	}
	return &request, nil
}

// GenerateInitHash generates the init hash for the drain contract
func GenerateInitHash(owner string) ([]byte, []byte, error) {
	args, err := drainABI.Pack("", []any{
		common.HexToAddress(owner),
	}...)
	if err != nil {
		return nil, nil, fmt.Errorf("could not pack drain constructor arguments: %v", err)
	}
	data := append(drainBytecode, args...)

	return crypto.Keccak256(data), data, nil
}

// Derive the contract address
// addr: the address of the factory contract
// owner: the address where the funds will be sent
// salt: the salt used to derive the address
func derivePaymentAddress(addr, owner, saltStr string) (common.Address, string, error) {
	var salt [32]byte
	if saltStr == "" {
		salt = generateRandomSalt()
	} else {
		salt = saltToBytes32(saltStr)
	}

	hash, _, err := GenerateInitHash(owner)
	if err != nil {
		return common.Address{}, "", fmt.Errorf("could not pack drain constructor arguments: %v", err)
	}
	return crypto.CreateAddress2(common.HexToAddress(addr), salt, hash), saltToHex(salt), nil
}

func saltToBytes32(salt string) [32]byte {
	var b [32]byte
	copy(b[:], salt)
	return b
}

func saltToHex(salt [32]byte) string {
	return fmt.Sprintf("0x%x", salt[:])
}

func generateRandomSalt() [32]byte {
	n := rand.Reader
	var b [32]byte
	if _, err := n.Read(b[:]); err != nil {
		panic(err)
	}
	return b
}
