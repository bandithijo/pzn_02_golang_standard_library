package main

import (
	"fmt"
	"math"
)

/*
Package math
- Package math merupakan package yang berisikan constant dan fungsi matematika
- <https://pkg.go.dev/math>

Beberapa Function di Package math
- math.Round(float64)
  Membulatkan float64 ke atas atau ke bawah, sesuai dengan yang paling dekat
- math.Floor(float64)
  Membulatkan float64 ke bawah
- math.Ceil(float64)
  Membulatkan float64 ke atas
- math.Max(float64, float64)
  Mengembalikan nilai float64 paling besar
- math.Min(float64, float64)
  Mengembalikan nilai float64 paling kecil
*/

func main() {
	fmt.Println(math.Ceil(1.40)) // => 2
	fmt.Println(math.Floor(1.60)) // => 1
	fmt.Println(math.Round(1.60)) // => 2
	fmt.Println(math.Max(10, 11)) // => 11
	fmt.Println(math.Min(10, 11)) // => 10
}
