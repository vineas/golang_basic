package main

import "fmt"

func main() {
	//map adalah kumpulan key-value, berguna untuk membuat data map berdasarkan key dan value masing-masing
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0
	// Kenapa mei = 0? Karena mei tidak ada di map chicken, sehingga nilainya adalah 0.

	//Inisialisasi map
	var data map[string]int
	// data["one"] = 1
	// akan muncul error!

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara vertikal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara horizontal
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken1, chicken2)

	/* variabel map bisa di inisialisasikan dengan tanpa nilai awal dengan menggunakan {} */
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	/* Iterasi Item Map menggunakan for - range */
	var chicken4 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken4 {
		fmt.Println(key, "  \t:", val)
	}

	/* Menghapus Item Map */
	var chicken5 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken5)) // 2
	fmt.Println(chicken5)

	delete(chicken5, "januari")

	fmt.Println(len(chicken5)) // 1
	fmt.Println(chicken5)

	/* Memeriksa Keberadaan Item Map */
	var value, isExist = chicken5["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	/* Kombinasi Slice dan Map */
	var chicken7 = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, val := range chicken7 {
		fmt.Println(val["gender"], val["name"])
	}

	/*
		Bisa dipersingkat dengan go terbaru:
		var chickens = []map[string]string{
			{"name": "chicken blue",   "gender": "male"},
			{"name": "chicken red",    "gender": "male"},
			{"name": "chicken yellow", "gender": "female"},
		}

		dalam []map[string]string, setiap elemen bisa saja memiliki key berbeda beda
		var data = []map[string]string{
			{"name": "chicken blue", "gender": "male", "color": "brown"},
			{"address": "mangga street", "id": "k001"},
			{"community": "chicken lovers"}
		}
	*/

}
