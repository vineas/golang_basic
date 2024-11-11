package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// momen seperti while
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	//Penggunaan keyword for tanpa argumen
	var x = 0
	for {
		fmt.Println("Angka", x)
		x++
		if x == 5 {
			break
		}
	}

	// break & continue, break memaksa untuk berhenti, continue memaksa untuk melanjutkan
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// perulangan memanfaatkan label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
