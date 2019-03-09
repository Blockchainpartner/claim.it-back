package db

import (
	"context"
	"github.com/Blockchainpartner/claim.it-back/util"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var (
	ClaimCollection mongo.Collection
	UserCollection  mongo.Collection
)

// initialize connection to the DB according to an environment variable
// for the DB URI
func Init() {
	/** PHASE 1: Init connection to the DB */
	// get DB URI from environment variable
	mongodbUri := util.MongodbUri
	// create client and init connection
	client, err := mongo.NewClient(mongodbUri)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// get DB name from environment variables
	dbName := util.MongodbDBName
	// create DB client object
	db := client.Database(dbName)

	/** PHASE 2: Init access to collections */
	UserCollection = *db.Collection("user")
	ClaimCollection = *db.Collection("claim")
}
