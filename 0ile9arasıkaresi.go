package main

import "fmt"

func main() {
	/* a := []int{1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i]*a[i])
	} */

	myArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(square(myArr))

}

func square(arr []int) []int {
	/* for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	} */

	for i, v := range arr {
		arr[i] = v * v
	}
	return arr
}
