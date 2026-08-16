/*
	selain if expression untuk melakukan percabangan kita bisa juga menggunakan switch expression
	nah switch expression sangan sederhana di bandingkan dengan if
	biasa nya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

	sama seperti di if switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

	di go switch bisa di gunakan tanpa kondisi
	jadi kondisi di switch tidak wajib
	jadi jika kita tidak menggunakan kondisi di switch expression,kita bisa menambahkan kondisi tersebut
	di setiap case nya

*/

package main

import "fmt"

func main() {
	name := "eko"

	switch name {
	case "eko":
		fmt.Println("hello eko")
	case "cecep":
		fmt.Println("hello cecep")
	default:
		fmt.Println("kamu bukan cecep atau eko")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("nama terlalu panjang")
	case length == 5:
		fmt.Println("nama sudah benar")
	case length< 5:
		fmt.Println("nama terlalu pendek")
	}
}