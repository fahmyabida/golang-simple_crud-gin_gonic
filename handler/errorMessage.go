package handler

import "github.com/gin-gonic/gin"

func ErrorQuery(err *error, result *gin.H) *gin.H {
	result = &gin.H{
		"message": "error query",
		"detail":  err,
	}
	return result
}

func ErrorScanQuery(err *error, result *gin.H) *gin.H {
	result = &gin.H{
		"message": "error scan query",
		"detail":  err,
	}
	return result
}
