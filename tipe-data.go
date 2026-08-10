/* Tipe Data di Go */
/* 
	Tipe data di Go ada 2 yaitu integer dan floating point.

	di tipe data integer ada beberapa jenis
	int8	nilai minimumnya -128 dan nilai maksimumnya 127
	int16	milai minimumnya -32768 dan nilai maksimumnya 32767
	int32 	nilai minimumnya -2147483648 dan nilai maksimumnya 2147483647
	int64	nilai minimumnya -9223372036854775808 dan nilai maksimumnya 9223372036854775807

	uint8	nilai minimumnya 0 dan nilai maksimumnya 255
	uint16	nilai minimumnya 0 dan nilai maksimumnya 65535
	uint32	nilai minimumnya 0 dan nilai maksimumnya 4294967295
	uint64	nilai minimumnya 0 dan nilai maksimumnya 18446744073709551615

	di tipe data floating point ada beberapa jenis
	float32 nilai minimumnya 1.18x10^-38 dan nilai maksimumnya 3.4x10^38
	float64 nilai minimumnya 2.23x10^-308 dan nilai maksimumnya 1.80x10^308
	complex64 nomer kompleks dengan float32 nyata dan bagian imajiner
	complex128 nomer kompleks dengan float64 nyata dan bagian imajiner

	dan ada nama alias untuk tipe data integer
	byte alias untuk uint8
	rune alias untuk int32
	int alias untuk int32 atau int 64 tergantung dari sistem oprerasi yang dipakai
	uint alias untuk uint32 atau uint64 tergantung dari sistem operasi yang dipakai
*/

package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
}