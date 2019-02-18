package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"halo": "it must return All data heros"})
}
