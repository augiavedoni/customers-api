package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func validateId(id string, context *gin.Context) bool {
	if id == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID must not be empty"})
		return false
	}

	int1, err := strconv.ParseInt(id, 6, 12)

	if err != nil || int1 < 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID is not valid"})
		return false
	} else {
		return true
	}
}
