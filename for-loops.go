/*
	di go lang hanya ada satu perulangan yaitu for loops,yang biasanya
	di bahasa pemograman lain ada 3 yaitu for while dan do while

	di dalam for loops itu kita bisa menambahkan statement dimana ada 2 statement
	yang bisa di tambahkan
	init statement itu statement sebelum for di eksekusi
	dan post statement itu statement yang akan selalu dieksekusi di akhir tiap perulangan

	for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	data collection contohnya array slice dan map
*/

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	// for counter2 := 1; counter2 <= 10; counter2++ {
	// 	fmt.Println("perulangan ke ", counter2)
	// }

	slice := []string{"deryl","fabiensyah","cihuyy"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"deryl","fabiensyah","cihuyyy"}
	for index, name := range names {
		fmt.Println("index",index, "=",name)
	}
}