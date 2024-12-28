package handler

import (
	//"errors"
	"log"
	"net/http"

	//secerr "github.com/Lsortudo/secretF/internal/errors"
	"github.com/Lsortudo/secretF/internal/errors"
	"github.com/Lsortudo/secretF/internal/service"
	"github.com/gin-gonic/gin"
)

// Storage for codes and their pairs
var storage = make(map[string][]service.Pair)

// InitializeStorage initializes the storage map
func InitializeStorage() {
	storage = make(map[string][]service.Pair)
}

// RequestPayload represents the expected JSON payload
type RequestPayload struct {
	Code   string   `json:"code"`
	People []string `json:"people"`
}

// DrawSecretSanta handles the POST /draw endpoint
func DrawSecretSanta(c *gin.Context) {
	var payload RequestPayload

	// Bind JSON input to payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidJSON.Error()})
		return
	}

	// Validate the code
	if len(payload.Code) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code must be exactly 5 characters"})
		return
	}

	// Generate pairs
	pairs, err := service.DrawPairs(payload.People)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Store the pairs with the code
	storage[payload.Code] = pairs

	// Return the result
	c.JSON(http.StatusOK, gin.H{
		"code":  payload.Code,
		"pairs": pairs,
	})
}

// VerifySecretSanta handles the GET /verify/:code endpoint
func VerifySecretSanta(c *gin.Context) {
	code := c.Param("code")

	// Check if the code exists in storage
	pairs, exists := storage[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "code not found"})
		return
	}

	// Return the pairs
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"pairs": pairs,
	})
}
