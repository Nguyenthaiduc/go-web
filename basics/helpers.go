package main

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano()) //random
	value := rand.Intn(n)

	return value
}
