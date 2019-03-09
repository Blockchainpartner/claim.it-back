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
	PusherAppId      string
	PusherAppKey     string
	PusherAppSecret  string
	PusherAppCluster string
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

	PusherAppId, exists = os.LookupEnv("PUSHER_APP_ID")
	if !exists {
		log.Fatal("PUSHER_APP_ID environment variable not set")
	}

	PusherAppKey, exists = os.LookupEnv("PUSHER_APP_KEY")
	if !exists {
		log.Fatal("PUSHER_APP_KEY environment variable not set")
	}

	PusherAppSecret, exists = os.LookupEnv("PUSHER_APP_SECRET")
	if !exists {
		log.Fatal("PUSHER_APP_SECRET environment variable not set")
	}

	PusherAppCluster, exists = os.LookupEnv("PUSHER_APP_CLUSTER")
	if !exists {
		log.Fatal("PUSHER_APP_CLUSTER environment variable not set")
	}
}
