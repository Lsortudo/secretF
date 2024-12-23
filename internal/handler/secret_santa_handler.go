package handler

import (
	"errors"
	"log"
	"net/http"

	secerr "github.com/Lsortudo/secretF/internal/errors"
	"github.com/Lsortudo/secretF/internal/service"
	"github.com/gin-gonic/gin"
)

func DrawSecretSanta(c *gin.Context) {
	log.Println("Received request at /draw")

	var people []string
	if err := c.ShouldBindJSON(&people); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": secerr.ErrInvalidJSON.Error()})
		return
	}

	log.Println("List of participants:", people)

	// Draw the pairs
	pairs, err := service.DrawPairs(people)
	if err != nil {
		log.Println("Error drawing pairs:", err)
		if errors.Is(err, secerr.ErrOddNumberOfPeople) { // Comparação correta
			c.JSON(http.StatusBadRequest, gin.H{"error": secerr.ErrOddNumberOfPeople.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": secerr.ErrInternalError.Error()})
		}
		return
	}

	log.Println("Drawn pairs:", pairs)
	c.JSON(http.StatusOK, pairs)
}
