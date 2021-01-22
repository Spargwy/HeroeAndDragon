package main

import (
	"math/rand"
	"time"
)

func randomSeed() int {
	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(100)
	return seed
}
