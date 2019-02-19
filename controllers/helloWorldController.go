package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Welcome": "Hello World!"})
}
