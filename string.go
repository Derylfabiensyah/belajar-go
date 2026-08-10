/* 
	tipe data string itu sekumpulan karaktrer yang jumlah nya bisa nol atau tak terhingga
	tipe data string di Go dipresentasikan dengan kata kunci string
	nilai string di Go selalu diawali dengan peting dua (") dan diakhiri dengan peting dua (")

	ada 2 function untuk String
	len("string") untuk menghitung jumlah karakter di String
	"string"[number] untuk mengambl karakter pada posisi yang ditentukan
*/

package main

import "fmt"

func main() {
	fmt.Println(len("Muhammad Deryl Fabiensyah"))
	fmt.Println("Nama Awal = Muhammad"[0])
	fmt.Println("Nama Tengah = Deryl"[1])
	fmt.Println("Nama Akhir = Fabiensyah"[2])

}