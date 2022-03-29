package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Notunuzu giriniz:")
	grade, _ := getGrade()

	if grade >= 50 {
		fmt.Println("geçtin")
	} else {
		fmt.Println("kaldın")
	}

}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}
