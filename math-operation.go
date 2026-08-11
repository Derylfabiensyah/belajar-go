/*
	operasi metematikan sama seperti bahasa pemograman lain
	ada + - * / %

	ada juga augmented assignment operator seperti
	a += 10 jika di operasi matematika a = a + 10
	a -= 10 jika di operasi matematika a = a - 10
	a *= 10 jika di operasi matematika a = a * 10
	a /= 10 jika di operasi matematika a = a / 10
	a %= 10 jika di operasi matematika a = a % 10

	ada juga unary operator seperti
	++ itu a = a + 1
	-- itu a = a - 1
	- itu negative number
	+ itu positive number tapi secara default number itu positif jadi tidak wajib pakai ini
	! itu untuk boolean kebalikan 
*/

package main

import "fmt"

func main() {
	var a = 10
	var b = 10

	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	var i = 10
	i += 10

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(i)
}