package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(catAndMouse(7, 3, 5))
}

func catAndMouse(x int32, y int32, z int32) string {
	x1 := math.Abs(float64(x) - float64(z))
	y1 := math.Abs(float64(y) - float64(z))
	if x1 < y1 {
		return "Cat A"
	} else if y1 < x1 {
		return "Cat B"
	} else {
		return "Mouse C"
	}

}
