/*
	if adalah salah satau kata kunci yang digunakan untuk percabangan
	percabangan artinya kita bisa mengeksekusi kode program tertentu
	kertika suatu kondisi terpenuhi
	hampir di semua bahasa pemograman mendukung if expression

	blok if akan dieksekusi ketika kondisi if bernilai true
	kadang kita ingin melakukan eksekusi program tertentu jika
	kondisi if bernilai false
	hal ini dilakukan menggunakan else expression

	kadang kadang dalam if kita butuh membukat beberapa kondisi
	kasus seperti ini kita bisa menggunakan Else If expression

	di go lang ada istilah if dengan short statement
	jadi if mendukung short statement sebelum kondisi
	hali ini sangat cocok untuk membuat statement yang sederhana 
	sebelum melakukan pengecekan terhadap kondisi
*/

package main

import "fmt"

func main() {
	name := "budi"

	if name == "eko" {
		fmt.Println("hello eko")
	}else if name == "cecep" {
		fmt.Println("hello cecep")
	}else {
		fmt.Println("kamu bukan eko atau pun cecep")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("nama sudah benar")
	}
}