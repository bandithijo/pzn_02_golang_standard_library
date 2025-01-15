package main

import (
	"fmt"
	"slices"
)

/*
Package slices
- Di Go versi terbaru, terdapat fitur bernama Generic, fitur ini akan kita bahas khusus di kelas Golang Generic
- Fitur Generic ini membuat kita bisa membuat parameter dengan tipe yang bisa berubah-ubah tanpa harus menggunakan interface kosong atau any
- Salah satu package yang menggunakan fitur Generic ini adalah package slices
- Package slices ini digunakan untuk memanipulasi data slice
- <https://pkg.go.dev/slices>
*/

func main() {
	var names []string = []string{"John", "Paul", "George", "Ringo"}
	var values []int = []int{100, 95, 80, 90}

	fmt.Println(slices.Min(values)) // => 80
	fmt.Println(slices.Max(values)) // => 100
	fmt.Println(slices.Min(names)) // => George
	fmt.Println(slices.Max(names)) // => Ringo
	fmt.Println(slices.Contains(names, "Rizqi")) // => false
	fmt.Println(slices.Index(names, "Rizqi")) // => -1
	fmt.Println(slices.Index(names, "George")) // => 2
}
