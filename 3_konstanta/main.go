package main

import "fmt"

func main() {
	// Konstanta
	const phi = 3.14
	fmt.Println(phi)

	// var dapat diubah nilainya
	var name string = "rapunzel"
	name = "john wick"
	fmt.Println(name)

	const firstName string = "john"
	// firstName = "john wick" //todak bisa diubah karna sudah konstan
	fmt.Print("halo ", firstName, "!\n")

	// perbedaan Print dan Println

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}
