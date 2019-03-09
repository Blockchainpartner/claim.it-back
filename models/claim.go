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

type Claim struct {
	ID            *objectid.ObjectID `json:"id" bson:"_id,omitempty"`
	TxHash        *string            `json:"txHash" bson:"txHash"`
	ClaimId       *string            `json:"claimId" bson:"claimId"`
	ParentClaimId *string            `json:"parentClaimId" bson:"parentClaimId"`
	Approved      *bool              `json:"approved" bson:"approved"`
	Valid         *bool              `json:"valid" bson:"valid"`
	Topic         *string            `json:"topic" bson:"topic"`
	Issuer        *string            `json:"issuer" bson:"issuer"`
	Recipient     *string            `json:"recipient" bson:"recipient"`
	Data          *string            `json:"data" bson:"data"`
	DataHash      *string            `json:"dataHash" bson:"dataHash"`
}

type ClaimFilter struct {
	ID            *objectid.ObjectID `json:"id" bson:"_id,omitempty,omitempty"`
	TxHash        *string            `json:"txHash,omitempty" bson:"txHash,omitempty"`
	ClaimId       *string            `json:"claimId,omitempty" bson:"claimId,omitempty"`
	ParentClaimId *string            `json:"parentClaimId,omitempty" bson:"parentClaimId,omitempty"`
	Approved      *bool              `json:"approved,omitempty" bson:"approved,omitempty"`
	Valid         *bool              `json:"valid,omitempty" bson:"valid,omitempty"`
	Topic         *string            `json:"topic,omitempty" bson:"topic,omitempty"`
	Issuer        *string            `json:"issuer,omitempty" bson:"issuer,omitempty"`
	Recipient     *string            `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Data          *string            `json:"data,omitempty" bson:"data,omitempty"`
	DataHash      *string            `json:"dataHash,omitempty" bson:"dataHash,omitempty"`
}

type ClaimUpdate map[string]ClaimFilter

func (u Claim) Post() error {
	// get claim collection
	claimColl := db.ClaimCollection

	// store claim in the db
	res, err := claimColl.InsertOne(
		context.Background(),
		u,
	)
	if err != nil {
		log.Println("Insert claim error: ", err)
		return err
	}

	// return
	log.Println("Inserted claim successfully: ", res)
	return nil
}

func (u Claim) Put(claimUpdate ClaimUpdate, options ...updateopt.Update) error {
	// get claim collection
	claimColl := db.ClaimCollection

	// update in the db
	res, err := claimColl.UpdateOne(
		context.Background(),
		ClaimFilter{ID: u.ID},
		claimUpdate,
		options...,
	)
	if err != nil {
		log.Println("Update claim error: ", err)
		return err
	}
	log.Println("Updated claim successfully: ", res)

	return nil
}

func GetClaim(claimFilter ClaimFilter) (*Claim, error) {
	// get claim collection
	claimColl := db.ClaimCollection
	// find document in DB
	doc := claimColl.FindOne(
		context.Background(),
		claimFilter,
	)
	// decode and return
	item := Claim{}
	err := doc.Decode(&item)
	if err != nil {
		log.Println("Decode error ", err)
		return &item, err
	}
	return &item, nil
}

func SearchClaims(rf RequestFilter) ([]ClaimFilter, error) {
	// get filter
	filter := rf.Filter

	// get find options
	findOptions := new(findopt.FindBundle).
		Projection(rf.Projection).
		Sort(rf.Sort).
		Skip(rf.Skip).
		Limit(rf.Limit)

	// get documents from DB
	claimColl := db.ClaimCollection
	cur, err := claimColl.Find(
		context.Background(),
		filter,
		findOptions,
	)
	if err != nil {
		return nil, err
	}

	// iterate through cursor to create return object
	defer cur.Close(context.Background())
	claims := make([]ClaimFilter, 0)
	for cur.Next(context.Background()) {
		// decode and return
		claim := ClaimFilter{}
		err := cur.Decode(&claim)
		if err != nil {
			return nil, err
		}
		claims = append(claims, claim)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	// return
	return claims, nil
}

func CountClaims(rf RequestFilter) (int64, error) {
	// get filter
	filter := rf.Filter

	// get find options
	countOptions := new(countopt.CountBundle).
		Skip(rf.Skip).
		Limit(rf.Limit)

	// get documents from DB
	claimColl := db.ClaimCollection
	nbDocs, err := claimColl.Count(
		context.Background(),
		filter,
		countOptions,
	)
	if err != nil {
		return 0, err
	}
	return nbDocs, nil
}

func (u Claim) Delete() error {
	// get claimSecret collection
	claimColl := db.ClaimCollection
	// delete document in DB and return
	_, err := claimColl.DeleteOne(
		context.Background(),
		u,
	)
	return err
}
