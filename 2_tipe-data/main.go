package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89                        //uint untuk bilangan positif saja
	var negativeNumber = -1243423644                     //int untuk bilangan bulat (negatif sampai positif)
	fmt.Printf("bilangan positif: %d\n", positiveNumber) //%d untuk bilangan bulat
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)   //%f untuk bilangan desimal
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) //menampilkan 3 angka dibelakang koma

	var exist bool = true
	fmt.Printf("exist? %t \n", exist) //%t untuk boolean

	var message string = "Halo"
	fmt.Printf("message: %s \n", message) //%s untuk string

	var messageStr = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`
	// Keistimewaan string menggunakan backtick adalah tidak di escape
	fmt.Println(messageStr)

}
