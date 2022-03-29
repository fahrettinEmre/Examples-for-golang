package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Merhaba\nŞanslı sayı :", randomSayı(2))
}

func randomSayı(seedValue int64) int {
	rand.Seed(seedValue)
	var sanslıSayı int
	sanslıSayı = rand.Intn(100)
	return sanslıSayı
}
