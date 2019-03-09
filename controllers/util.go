package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UtilController struct{}

func (uc UtilController) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}
