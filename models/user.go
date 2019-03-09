package models

import (
	"context"
	"github.com/Blockchainpartner/claim.it-back/db"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo/countopt"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
	"log"
)

type User struct {
	ID                 *objectid.ObjectID `json:"id" bson:"_id,omitempty"`
	ProxyAddress       *string            `json:"proxyAddress" bson:"proxyAddress"`
	ActionKeyAddress   *string            `json:"actionKeyAddress" bson:"actionKeyAddress"`
	ActionKeyPublicKey *string            `json:"actionKeyPublicKey" bson:"actionKeyPublicKey"`
	ClaimHolderAddress *string            `json:"claimHolderAddress" bson:"claimHolderAddress"`
	Pseudonym          *string            `json:"pseudonym" bson:"pseudonym"`
	PictureUri         *string            `json:"pictureUri" bson:"pictureUri"`
}

type UserFilter struct {
	ID                 *objectid.ObjectID `json:"id" bson:"_id,omitempty"`
	ProxyAddress       *string            `json:"proxyAddress,omitempty" bson:"proxyAddress,omitempty"`
	ActionKeyAddress   *string            `json:"actionKeyAddress,omitempty" bson:"actionKeyAddress,omitempty"`
	ActionKeyPublicKey *string            `json:"actionKeyPublicKey,omitempty" bson:"actionKeyPublicKey,omitempty"`
	ClaimHolderAddress *string            `json:"claimHolderAddress,omitempty" bson:"claimHolderAddress,omitempty"`
	Pseudonym          *string            `json:"pseudonym,omitempty" bson:"pseudonym,omitempty"`
	PictureUri         *string            `json:"pictureUri,omitempty" bson:"pictureUri,omitempty"`
}

type UserUpdate map[string]UserFilter

func (u User) Post() error {
	// get user collection
	userColl := db.UserCollection

	// store user in the db
	res, err := userColl.InsertOne(
		context.Background(),
		u,
	)
	if err != nil {
		log.Println("Insert user error: ", err)
		return err
	}

	// return
	log.Println("Inserted user successfully: ", res)
	return nil
}

func (u User) Put(userUpdate UserUpdate, options ...updateopt.Update) error {
	// get user collection
	userColl := db.UserCollection

	// update in the db
	res, err := userColl.UpdateOne(
		context.Background(),
		UserFilter{ID: u.ID},
		userUpdate,
		options...,
	)
	if err != nil {
		log.Println("Update user error: ", err)
		return err
	}
	log.Println("Updated user successfully: ", res)

	return nil
}

func GetUser(userFilter UserFilter) (*User, error) {
	// get user collection
	userColl := db.UserCollection
	// find document in DB
	doc := userColl.FindOne(
		context.Background(),
		userFilter,
	)
	// decode and return
	item := User{}
	err := doc.Decode(&item)
	if err != nil {
		log.Println("Decode error ", err)
		return &item, err
	}
	return &item, nil
}

func SearchUsers(rf RequestFilter) ([]UserFilter, error) {
	// get filter
	// TODO make filter on address case insensitive
	filter := rf.Filter

	// get find options
	findOptions := new(findopt.FindBundle).
		Projection(rf.Projection).
		Sort(rf.Sort).
		Skip(rf.Skip).
		Limit(rf.Limit)

	// get documents from DB
	userColl := db.UserCollection
	cur, err := userColl.Find(
		context.Background(),
		filter,
		findOptions,
	)
	if err != nil {
		return nil, err
	}

	// iterate through cursor to create return object
	defer cur.Close(context.Background())
	users := make([]UserFilter, 0)
	for cur.Next(context.Background()) {
		// decode and return
		user := UserFilter{}
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	// return
	return users, nil
}

func CountUsers(rf RequestFilter) (int64, error) {
	// get filter
	filter := rf.Filter

	// get find options
	countOptions := new(countopt.CountBundle).
		Skip(rf.Skip).
		Limit(rf.Limit)

	// get documents from DB
	userColl := db.UserCollection
	nbDocs, err := userColl.Count(
		context.Background(),
		filter,
		countOptions,
	)
	if err != nil {
		return 0, err
	}
	return nbDocs, nil
}

func (u User) Delete() error {
	// get userSecret collection
	userColl := db.UserCollection
	// delete document in DB and return
	_, err := userColl.DeleteOne(
		context.Background(),
		u,
	)
	return err
}
