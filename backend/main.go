package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func main() {
	if err := LoadConfig("."); err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	db, err := LoadDatabase()
	if err != nil {
		log.Fatalf("Error loading database: %s", err)
	}

	// Parse private key
	pk, err := parsePrivateKey(viper.GetString("factory_pk")[2:])
	if err != nil {
		log.Fatalf("Error parsing private key: %s", err)
	}

	// Connect to Ethereum node
	client, err := ethclient.Dial(viper.GetString("eth_node"))
	if err != nil {
		log.Fatalf("Error connecting to Ethereum node: %s", err)
	}

	// Run scheduler
	go RunScheduler(time.Second*10, client, pk, db)

	fmt.Println("Starting server on port", viper.GetString("listen_addr"))
	server := &http.Server{
		Addr:         viper.GetString("listen_addr"),
		Handler:      NewRouter(db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
