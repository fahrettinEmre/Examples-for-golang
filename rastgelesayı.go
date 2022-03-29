package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)
	fmt.Println("1 ile 100 arasındaki sayıyı tahmin edin")
	reader := bufio.NewReader(os.Stdin)

	for attemts := 0; attemts < 10; attemts++ {
		fmt.Println(10-attemts, "hakkınız var")
		fmt.Println("Lütfen tahmin yapınız")

		input, _ := reader.ReadString('\n')
		//if err != nil {
		//	fmt.Println(err)
		//}

		input = strings.TrimSpace(input)
		num, _ := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Tahmin büyük, daha küçük sayı giriniz")
		} else if num < target {
			fmt.Println("Tahmin küçük, daha büyük sayı giriniz")
		} else {
			fmt.Println("Doğru tahmin hedef", target, attemts, "seferde buldunuz")
			break
		}

	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
