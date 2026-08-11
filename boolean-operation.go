/*
	jadi operasi boolean adalah operasi yang hanya dilakukan di data boolean saja
	ada 3 operasi boolean yaitu
	&& sebagai dan
	|| sebagai atau
	! sebagai kebalikan

	jadi operasi && itu jika kedua nilai berisi true maka hasilnya akan true juga,
	tapi jika salah satu nilai berisi false maka hasilnya akan false

	jadi operasi || itu jika kedua nilai nya berisi true maka hasilnya akan true,
	dan jika ada salah satu nilai yang berisikan false maka hasilnya akan true juga,
	tapi jika kedua nilai nya false hasilnya akan false juga

	nah kalo ! atau kebalikan itu ketika nilainya true maka hasilnya akan false,
	tapi jika nilai nya false maka hasilnya akan true


*/

package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}