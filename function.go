/*
	sebelumnya kira sudah mengenal sebuah function yang 
	wajib di buat agar program kita bisa berjalan yaitu function main
	function adalah sebuah blok kode yang sengaja di buat dalam program
	agar kita bisa digunakan berulang ulang
	cara membuat function sangat sederhana hanya dengan menggunakan kata kunci
	func lalu di ikuti dengna nama function nya dan blok kode isi function nya
	setelah membuat function kita bisa mengeksekusi function tersebut dengan 
	memanggilnya menggunakan kata kunci nama function nya di ikuti tanda kurung 
	buka dan kurung tutup
*/

package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func main() {
	sayHello()
}