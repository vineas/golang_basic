package main

import "fmt"

func main() {
	var firstName string = "john" //deklarasi dengan tipe data
	// var lastName string //ini kelebihan dari deklarasi dengan tipe data
	// lastName = "wick"

	//deklarasi tanpa tipe data
	lastName := "wick"

	// deklarasi multiple variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	// lebih ringkas
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// variabel underscore (_) untuk menampung nilai yang tidak dipakai/keranjang sampah
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	// name, _ := "john", "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println(first, second, third)
	fmt.Println(seventh, eight, ninth)
	//%s disini sebagai string yang akan ditampilkan
	//sementara \n disini sebagai enter
	// printf digunakan untuk menampilkan string, sementara println digunakan untuk enter

	name := new(string)
	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

}
