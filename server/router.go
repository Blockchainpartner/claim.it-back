package server

import (
	"github.com/Blockchainpartner/claim.it-back/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// create a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// declare endpoints
	utilGroup := router.Group("/util")
	{
		uc := controllers.UtilController{}
		utilGroup.GET("/health", uc.GetHealth)
	}

	userGroup := router.Group("/user")
	{
		uc := controllers.UserController{}
		userGroup.POST("", uc.PostUser)
		userGroup.POST("/search", uc.SearchUser)
		userGroup.POST("/search/count", uc.CountUser)
		userGroup.GET("/:id", uc.GetUser)
		userGroup.PUT("/:id", uc.PutUser)
	}

	claimGroup := router.Group("/claim")
	{
		uc := controllers.ClaimController{}
		claimGroup.POST("", uc.PostClaim)
		claimGroup.POST("/search", uc.SearchClaim)
		claimGroup.POST("/search/count", uc.CountClaim)
		claimGroup.GET("/:id", uc.GetClaim)
		claimGroup.PUT("/:id", uc.PutClaim)
	}

	return router
}
