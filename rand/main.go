package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min, max := int64(0), int64(1)
	// [min, max)
	log.Println(rand.Int63n(max-min) + min)
	// [min, max]
	log.Println(rand.Int63n(max-min+1) + min)
}
