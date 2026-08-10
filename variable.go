/* 
	variable adalah tempat menyimpan data dan variable digunakan agar kita bisa
	akses data yang di simpan dimanapun kita mau
	di Go variable hanya nosa menyimpan tipe data yang sama,jika ingin mentimpan
	data yang berbeda beda jenis,kita harus memuat beberapa variable
	untuk membuat variable kira bisa menggunakan kata kunci var,lalu diikuti 
	dengan nama variable dan tipe datanya

	saat kita mmebuat variable kita wajib untuk menyebutkan tipe data dari variable tersebut,
	namun jika kita langsung menginisialisasikan sata pada variable nya, maka kita tidak wajib 
	menyebutkan tipe data variablenya

	di Go kata kunci var saat mmebuat variable tidak lah wajib
	asalkan saat membuat variable kita langsung menginisialisasi datanya agar tidak perlu menggunakan
	kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikandata pada variable tersebut
	kata kunci := hanya untuk deklarasi awal saja

	ada deklarasi multiple variable,jadi di Go kita bisa membuat beberapa variable sekaligus banyak 
	dengan cara tambah kan tanda kurung buka dan kurung tutup setelah kata kunci var, llau di dalam 
	
*/

package main

import "fmt"

func main() {
	var name string

	name = "Deryl Fabiensyah"
	fmt.Println(name)

	name = "Muhammad Deryl Fabiensyah"
	fmt.Println(name)

	var umur = 16
	fmt.Println(umur)

	alamat := "Jl.Siliwangin No. 46"
	fmt.Println(alamat)

	var (
		firstName = "Muhammad"
		lastName = "Fabiensyah"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}