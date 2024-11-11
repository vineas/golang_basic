package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// names[4] = "d" jika ditambah akan error
	fmt.Println(names[0], names[1], names[2], names[3])

	//inisialisasi nilai awal array *gaya horizontal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// *gaya vertikal
	var hewan = [4]string{
		"anjing",
		"kelinci",
		"ayam",
		"panda",
	}
	fmt.Println("Jumlah element \t\t", len(hewan))
	fmt.Println("Isi semua element \t", hewan)

	//array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// looping array menggunakan for
	var asiaCountry = [4]string{"indonesia", "malaysia", "singapura", "jepang"}
	for i := 0; i < len(asiaCountry); i++ {
		fmt.Printf("elemen %d: %s\n", i, asiaCountry[i])
	}

	// looping array menggunakan for range
	for i, country := range asiaCountry {
		fmt.Printf("elemen %d: %s\n", i, country)
	}

	// penggunaan variabel underscore _, dalam for - range
	for _, country := range asiaCountry {
		fmt.Printf("elemen: %s\n", country)
	}

	// alokasi array menggunakan keyword make
	var names2 = make([]string, 4)
	names2[0] = "trafalgar"
	names2[1] = "d"
	names2[2] = "water"
	names2[3] = "law"
	fmt.Println(names2)

}
