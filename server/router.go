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

	return router
}
