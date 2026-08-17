/*
	break dan continue adalah kata kunci yangb isa di gunakan dalam perulangan
	break digunakan untuk menghentikan seluruh perulangan
	continue adalah digunakan untuk menghentikan perulangan yang berjalan dan 
	langsung menalajutkan ke perulangan selanjutnya
*/

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan ke", i)
	}
}