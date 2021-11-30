package main

import (
	"math/rand"
	"time"
)

func randomNum() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(10)
}
