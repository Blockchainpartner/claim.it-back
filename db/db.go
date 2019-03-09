package db

import (
	"context"
	"github.com/Blockchainpartner/claim.it-back/util"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var UserCollection mongo.Collection

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

	/** PHASE 3: Init indexes */
	// Unique key on userCollection
	userCollectionIndexView := UserCollection.Indexes()
	_, err = userCollectionIndexView.CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.Int32("pseudonym", 1),
			),
			Options: mongo.NewIndexOptionsBuilder().
				Unique(true).
				Build(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
