package util

import (
	"log"
	"os"
)

var (
	MongodbDBName string
	MongodbUri    string
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
}
