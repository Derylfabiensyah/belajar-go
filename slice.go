/*
	tipe data slice itu potongan data dari array
	slice mirip dengan array yang membedakan ukuran slice bisa berubah
	slice dan array selalu terkoneksi dimana slice ada data yang mengakses sebagian atau seluruh 
	data di array

	di tipe data slice ada 3 data yaitu pointer adalah penunjuk data pertama di array para slice
	ada juga length adalah panjang dari slice
	dan yang terakhir ada capacity itu kapasitan dari slice,dimana length tidak boleh lebih dari capacity

	ada 4 cara membuat slice
	array[low:high]	membuat slice dari array mulai index low sampai index sebelum high
	array[low:]	membuat slice dari array dimulai dari index low sampai index akhir di array
	array[:high] membuat slice dari array di mulai dari index 0 sampai index sebelum high
	array[:] membuat slice dari array dimulai dari index 0 sampai index akhir di array
	
	di tipe data ini ada beberapa functon yaitu
	len(slice) untuk mendapatkan panjang
	cap(slice) untuk mendapatkan kapasitas
	append(slice, data) membuat slice baru dengan menambahkan data ke posisi terakhir slice,
	jika kapasitas sudah penuh maka akan membuat array baru
	make([]TypeData, length, capacity) untuk membuat slice baru
	copy(destination, source) untuk menyalin slice dari source ke destination
*/

package main

import "fmt"

func main() {
	names := [...]string{"Deryl","Kiano","Bama","Jibril","Fariz",}
	slice := names[4:5]

	fmt.Println(slice[0])
	fmt.Println(len(slice))
}