package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"} //tidak ditulis/diisi saat deklarasi variabel
	fmt.Println(fruits[0])                                     // "apple"

	// perbedaan slice dengan array adalah
	/*
		var fruitsA = []string{"apple", "grape"}     // slice
		var fruitsB = [2]string{"banana", "melon"}   // array
		var fruitsC = [...]string{"papaya", "grape"} // array
	*/

	// Hubungan slice dengan array & operasi slice
	var newFruits = fruits[1:3] //slice
	fmt.Println(newFruits)      //["grape", "banana"]

	// slice merupakan tipedata reference
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// Fungsi len
	fmt.Println(len(fruits)) // 4

	// Fungsi cap, untuk menghitung lebar array
	fmt.Println(cap(fruits)) // 4
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// Fungsi append, untuk menambahkan elemen di akhir slice
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits)  // [apple pinnaple banana melon papaya]
	fmt.Println(cFruits) // [apple pinnaple banana melon papaya]

	//Fungsi copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	// Pengaksesan elemen slice dengan 3 index
	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	var fruits3 = fruits2[1:3:3]
	fmt.Println(fruits2) // [apple grape banana melon]
	fmt.Println(fruits3) // [grape banana]
}
