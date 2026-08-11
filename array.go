/*
	tipe data array itu adalah tipe data yang berisikan kumpulan data yang tipenya sama
	saat membuat array itu kita perlu menentukan jumlah data yang bisa di tampung oleh array
	daya tampung array tidak bisa bertambah setelah array di buat

	di Go bisa membuat array secara langsung tidak manual mendeklarasikan jumlah data nya 
	terus tipe datanya lalu masukan datanya satu persatu

	ada tiga function array yaitu
	len(array) untuk mendapatkan panjang array
	array[index] untuk mendapat data di posisi index
	array[index] = value untuk mengubah data di posisi index
*/

package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Deryl"
	names[2] = "Fabiensyah"

	var nilai = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(nilai)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(len(names))
	fmt.Println(nilai[2])
	nilai[2] = 100
	fmt.Println(nilai)
}