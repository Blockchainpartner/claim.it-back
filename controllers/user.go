package controllers

import (
	"context"
	"github.com/Blockchainpartner/claim.it-back/ethereum"
	"github.com/Blockchainpartner/claim.it-back/models"
	"github.com/Blockchainpartner/claim.it-back/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"log"
	"math/big"
	"net/http"
	"strings"
)

type UserController struct {
}

func (uc UserController) PostUser(c *gin.Context) {
	// get new user fields from request parameter
	var userFilter models.UserFilter
	if err := c.ShouldBind(&userFilter); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	// get action key
	actionKeyString := userFilter.ActionKeyAddress
	if actionKeyString == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}
	*actionKeyString = strings.TrimPrefix(*actionKeyString, "0x")
	actionKey := common.HexToAddress(*actionKeyString)

	// handle the nonce manually
	nonce, err := ethereum.Client.PendingNonceAt(context.Background(), ethereum.Transactor.From)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	bigIntNonce := new(big.Int)

	// deploy identity contract
	bigIntNonce.SetUint64(nonce)
	ethereum.Transactor.Nonce = bigIntNonce
	identityAddress, _, _, err := ethereum.DeployIdentity(ethereum.Transactor, ethereum.Client, actionKey)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}

	// manually increment the nonce
	nonce++
	bigIntNonce.SetUint64(nonce)
	ethereum.Transactor.Nonce = bigIntNonce

	// deploy claim holder contract
	claimHolderAddress, _, _, err := ethereum.DeployClaimHolder(ethereum.Transactor, ethereum.Client, identityAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}

	// create new user object and force some fields.
	// the Email field should be left empty for instance
	userId := objectid.New()
	identityAddressPtr := identityAddress.Hex()
	claimHolderAddressPtr := claimHolderAddress.Hex()
	user := models.User{
		ID:                 &userId,
		ActionKeyAddress:   userFilter.ActionKeyAddress,
		ActionKeyPublicKey: userFilter.ActionKeyPublicKey,
		Pseudonym:          userFilter.Pseudonym,
		PictureUri:         userFilter.PictureUri,
		ProxyAddress:       &identityAddressPtr,
		ClaimHolderAddress: &claimHolderAddressPtr,
	}

	// post new user to the DB
	if err := user.Post(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	// return
	returnUserResponse(c, userId)
}

func (uc UserController) GetUser(c *gin.Context) {
	// get user
	contextUserIdString := c.Param("id")
	contextUserId, err := objectid.FromHex(contextUserIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := models.GetUser(models.UserFilter{ID: &contextUserId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// return
	returnUserResponse(c, *user.ID)
}

func (uc UserController) PutUser(c *gin.Context) {
	// get user
	contextUserIdString := c.Param("id")
	contextUserId, err := objectid.FromHex(contextUserIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := models.GetUser(models.UserFilter{ID: &contextUserId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// get new user fields from request parameter
	var newUser models.UserFilter
	if err := c.ShouldBind(&newUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user features"})
		return
	}

	// update user with non-zero values from newUser.
	var userSetFilter models.UserFilter
	if util.IsNotZero(newUser.ProxyAddress) {
		userSetFilter.ProxyAddress = newUser.ProxyAddress
	}
	if util.IsNotZero(newUser.ActionKeyAddress) {
		userSetFilter.ActionKeyAddress = newUser.ActionKeyAddress
	}
	if util.IsNotZero(newUser.ActionKeyPublicKey) {
		userSetFilter.ActionKeyPublicKey = newUser.ActionKeyPublicKey
	}
	if util.IsNotZero(newUser.ClaimHolderAddress) {
		userSetFilter.ClaimHolderAddress = newUser.ClaimHolderAddress
	}
	if util.IsNotZero(newUser.Pseudonym) {
		userSetFilter.Pseudonym = newUser.Pseudonym
	}
	if util.IsNotZero(newUser.PictureUri) {
		userSetFilter.PictureUri = newUser.PictureUri
	}

	// put update to the DB
	if userSetFilter != (models.UserFilter{}) { // check for non-empty filter
		update := models.UserUpdate{"$set": userSetFilter}
		if err := user.Put(update); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user features"})
			return
		}
	}

	// return
	returnUserResponse(c, *user.ID)
}

func (uc UserController) DeleteUser(c *gin.Context) {
	// get user
	contextUserIdString := c.Param("id")
	contextUserId, err := objectid.FromHex(contextUserIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := models.GetUser(models.UserFilter{ID: &contextUserId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// delete from the DB
	if err := user.Delete(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	// return
	c.Status(http.StatusOK)
	return
}

func (uc UserController) SearchUser(c *gin.Context) {
	// get filter fields from request parameter
	var requestFilter models.RequestFilter
	if err := c.ShouldBind(&requestFilter); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter"})
		return
	}

	// get documents from the db
	users, err := models.SearchUsers(requestFilter)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid filter"})
		return
	}

	// return
	c.JSON(http.StatusOK, users)
}

func (uc UserController) CountUser(c *gin.Context) {
	// get filter fields from request parameter
	var requestFilter models.RequestFilter
	if err := c.ShouldBind(&requestFilter); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter"})
		return
	}

	// get nb documents from the db
	nbUsers, err := models.CountUsers(requestFilter)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid filter"})
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{"count": nbUsers})
}

func returnUserResponse(c *gin.Context, userId objectid.ObjectID) {
	if user, err := models.GetUser(models.UserFilter{ID: &userId}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch user"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
