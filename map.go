/*
	 pada array atau slice untuk mengakses data kita menggunakan index number dimulai dari 0
	map adalah tipe data lain yang berisikan kumpulan data yang sama namun kita bisa menetukan
	jenis tipe data index yang akan kita gunakan jadi sederhana nya map adalah tipe data kumpulan
	key value (kata kunci nilai) dimana kata kuncinya bersifat unik tidak boleh sama
	berbeda dengan arrau dam slice jumlah data yang kita masukkan ke dalam map boleh
	sebanyak banyaknya asalkan kata kunci nya berbeda jika kita gunakan kata kunci sama
	maka otomatis data sebelumnya akan di ganti dengan data baru

	di go ada function map itu ada
	len(nap) untuk mendapatkan jumlah data di map
	map[key] untuk mengambil data di map dengan key
	map[key] = value untuk mengubah data di map dengan key
	make(map[typeKey]TypeValue) untuk membuat map baru
	delete(map, key) untuk menghapus data di map dengan key
*/

package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{
		
	// } cara deklarasi panjang nya

	person := map[string]string{ //ini cara deklalari singkatnya
		"name" : "deryl",
		"address" : "cianjur",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar go lang"
	book["author"] = "deryl"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

	
}