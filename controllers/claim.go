package controllers

import (
	"github.com/Blockchainpartner/claim.it-back/models"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"log"
	"net/http"
	"time"
)

type ClaimController struct {
}

func (uc ClaimController) PostClaim(c *gin.Context) {
	// get new claim fields from request parameter
	var claimFilter models.ClaimFilter
	if err := c.ShouldBind(&claimFilter); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim"})
		return
	}

	// TODO TX @Thomas

	// create new claim object and force some fields.
	// the Email field should be left empty for instance
	claimId := objectid.New()
	truePtr := true
	falsePtr := false
	now := time.Now().Unix()
	claim := models.Claim{
		ID:            &claimId,
		ClaimId:       claimFilter.ClaimId,
		ParentClaimId: claimFilter.ParentClaimId,
		Approved:      &falsePtr,
		Valid:         &truePtr,
		Topic:         claimFilter.Topic,
		Issuer:        claimFilter.Issuer,
		Recipient:     claimFilter.Recipient,
		Data:          claimFilter.Data,
		DataHash:      claimFilter.DataHash,
		Date:          &now, // TODO use timestamp from the block at mining instead
	}

	// post new claim to the DB
	if err := claim.Post(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim"})
		return
	}

	// return
	returnClaimResponse(c, claimId)
}

func (uc ClaimController) GetClaim(c *gin.Context) {
	// get claim
	contextClaimIdString := c.Param("id")
	contextClaimId, err := objectid.FromHex(contextClaimIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim id"})
		return
	}
	claim, err := models.GetClaim(models.ClaimFilter{ID: &contextClaimId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim id"})
		return
	}

	// return
	returnClaimResponse(c, *claim.ID)
}

func (uc ClaimController) PutClaim(c *gin.Context) {
	// get claim
	contextClaimIdString := c.Param("txHash")
	contextClaimId, err := objectid.FromHex(contextClaimIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim id"})
		return
	}
	claim, err := models.GetClaim(models.ClaimFilter{ID: &contextClaimId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim id"})
		return
	}

	// get new claim fields from request parameter
	var newClaim models.ClaimFilter
	if err := c.ShouldBind(&newClaim); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim features"})
		return
	}

	// update claim with non-zero values from newClaim.
	// TODO @thomas see what fields we can update
	var claimSetFilter models.ClaimFilter

	// put update to the DB
	if claimSetFilter != (models.ClaimFilter{}) { // check for non-empty filter
		update := models.ClaimUpdate{"$set": claimSetFilter}
		if err := claim.Put(update); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update claim features"})
			return
		}
	}

	// return
	returnClaimResponse(c, *claim.ID)
}

func (uc ClaimController) DeleteClaim(c *gin.Context) {
	// get claim
	contextClaimIdString := c.Param("txHash")
	contextClaimId, err := objectid.FromHex(contextClaimIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim id"})
		return
	}
	claim, err := models.GetClaim(models.ClaimFilter{ID: &contextClaimId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid claim id"})
		return
	}

	// delete from the DB
	if err := claim.Delete(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete claim"})
		return
	}

	// return
	c.Status(http.StatusOK)
	return
}

func (uc ClaimController) SearchClaim(c *gin.Context) {
	// get filter fields from request parameter
	var requestFilter models.RequestFilter
	if err := c.ShouldBind(&requestFilter); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter"})
		return
	}

	// get documents from the db
	claims, err := models.SearchClaims(requestFilter)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid filter"})
		return
	}

	// return
	c.JSON(http.StatusOK, claims)
}

func (uc ClaimController) CountClaim(c *gin.Context) {
	// get filter fields from request parameter
	var requestFilter models.RequestFilter
	if err := c.ShouldBind(&requestFilter); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter"})
		return
	}

	// get nb documents from the db
	nbClaims, err := models.CountClaims(requestFilter)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid filter"})
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{"count": nbClaims})
}

func returnClaimResponse(c *gin.Context, claimId objectid.ObjectID) {
	if claim, err := models.GetClaim(models.ClaimFilter{ID: &claimId}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch claim"})
	} else {
		c.JSON(http.StatusOK, claim)
	}
}
