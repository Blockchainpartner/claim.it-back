package util

import (
	"log"
	"os"
)

var (
	MongodbDBName    string
	MongodbUri       string
	EthNodeUri       string
	EthNodeWSUri     string
	HexEthPrivateKey string
)

func InitEnvironment() {
	var exists bool
	MongodbUri, exists = os.LookupEnv("MONGODB_URI")
	if !exists {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	MongodbDBName, exists = os.LookupEnv("MONGODB_DB_NAME")
	if !exists {
		log.Fatal("MONGODB_DB_NAME environment variable not set")
	}

	EthNodeUri, exists = os.LookupEnv("ETH_NODE_URI")
	if !exists {
		log.Fatal("ETH_NODE_URI environment variable not set")
	}

	EthNodeWSUri, exists = os.LookupEnv("ETH_NODE_WS_URI")
	if !exists {
		log.Fatal("ETH_NODE_URI environment variable not set")
	}

	HexEthPrivateKey, exists = os.LookupEnv("ETH_PRIVATE_KEY")
	if !exists {
		log.Fatal("ETH_PRIVATE_KEY environment variable not set")
	}
}
