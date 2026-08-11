/*
	type declaration itu adalah membuat ulang tipe data baru dari tipe data yang sudah ada
	type declaration itu biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada 
*/

package main

import "fmt"

func main() {
	type noKTP string
	type keterangan bool

	var noKtpDeryl noKTP = "1234567890"
	var lulusStatus keterangan = true
	fmt.Println(noKtpDeryl)
	fmt.Println(lulusStatus)
}