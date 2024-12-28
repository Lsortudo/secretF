package service

import (
	"math/rand"

	"github.com/Lsortudo/secretF/internal/errors"
)

// Pair represents a pair of participants
type Pair struct {
	Pair1 string `json:"pair1"`
	Pair2 string `json:"pair2"`
}

// DrawPairs generates random pairs from a list of participants
func DrawPairs(participants []string) ([]Pair, error) {
	if len(participants)%2 != 0 {
		return nil, errors.ErrOddNumberOfPeople
	}

	// Shuffle the participants list
	rand.Shuffle(len(participants), func(i, j int) {
		participants[i], participants[j] = participants[j], participants[i]
	})

	// Create pairs
	var pairs []Pair
	for i := 0; i < len(participants); i += 2 {
		pairs = append(pairs, Pair{
			Pair1: participants[i],
			Pair2: participants[i+1],
		})
	}

	return pairs, nil
}
