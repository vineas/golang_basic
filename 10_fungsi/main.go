package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"John", "Wick"}
	printMessage("halo", names)

	/*Fungsi dengan return Vaue */
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	/* Penggunaan fungsi rand.send() */
	// rand.Seed(time.Now().Unix())

	/* Import banyak Package */

	// import "fmt"
	// import "math/rand"
	// import "time"

	// // atau

	// import (
	//     "fmt"
	//     "math/rand"
	//     "time"
	// )

	/* Deklarasi parameter bertipe data sama */
	// func nameOfFunc(paramA type, paramB type, paramC type) returnType
	// func nameOfFunc(paramA, paramB, paramC type) returnType

	// func randomWithRange(min int, max int) int
	// func randomWithRange(min, max int) int

	/* Penggunaan keyword return untuk Menghentikan Proses dalam fungsi */
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
