package service

import (
	"math/rand"
	"time"

	"github.com/Lsortudo/secretF/internal/errors"
	"github.com/Lsortudo/secretF/internal/model"
)

func DrawPairs(people []string) ([]model.Sorteio, error) {
	if len(people)%2 != 0 {
		return nil, errors.ErrOddNumberOfPeople
	}

	// Shuffle the list of people randomly (coisa nova, tava fazendo manualmente kkkkkk)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})

	// Create pairs from the shuffled list
	var pairs []model.Sorteio
	for i := 0; i < len(people); i += 2 {
		pairs = append(pairs, model.Sorteio{
			Pair1: people[i],
			Pair2: people[i+1],
		})
	}

	return pairs, nil
}
