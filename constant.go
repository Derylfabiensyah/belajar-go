/* 
	constant adalah variable yang nilainya tidak bisa di ubah setelah pertama kali diberinilai
	cara pembuatan constant sama seperti variable yang membedakan hanya kata kunci yang di pakai
	kalo variable itu pakai var tapi kalo constant pakai const 
	saat pembuatan constant kita harus langsung menginisialisasikan nilainya 
	
	sama seperti variable constant juga bisa membuat multiple constant dengan cara menambahkan 
	tanda kurung buka dan kurung tutup setelah kata kunci const
*/

package main

import "fmt"

func main() {
	const firstName = "Muhammad"

	const (
		name = "Deryl"
		age = 16
	)

	/* firstName = "deryl"  error karena tidak bisa mengubah nilai constant*/
	fmt.Println(firstName)
	fmt.Println(name)
	fmt.Println(age)
}