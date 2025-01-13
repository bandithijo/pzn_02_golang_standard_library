package main

import (
	"fmt"
	"strconv"
)

/*
Package strconv (string conversion)
- Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int64
- Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
- Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
- <https://pkg.go.dev/strconv>

Beberapa Function di Package strconv
- strconv.parseBool(string)
  Mengubah string ke bool
- strconv.parseFloat(string)
  Mengubah string ke float
- strconv.parseInt(string)
  Mengubah string ke int64
- strconv.FormatBool(bool)
  Mengubah bool ke string
- strconv.FormatFloat(float, ...)
  Mengubah float64 ke string
- strconv.FormatInt(int, ...)
  Mengubah int64 ke string
*/

func errorHandling(value any, err error) {
	if err == nil {
		fmt.Println(value) // => true
	} else {
		fmt.Println("Error", err.Error())
	}
}

func main() {
	resultBool, err := strconv.ParseBool("benar")
	errorHandling(resultBool, err) // => Error strconv.ParseBool: parsing "bear": invalid syntax

	resultBool, err = strconv.ParseBool("true")
	errorHandling(resultBool, err) // => true

	resultInt, err := strconv.Atoi("lima ribu")
	errorHandling(resultInt, err) // => Error strconv.Atoi: parsing "lima ribu": invalid syntax

	resultInt, err = strconv.Atoi("5000")
	errorHandling(resultInt, err) // => 5000

	var binary string = strconv.FormatInt(5000, 2)
	fmt.Println(binary) // => "1001110001000"

	var resultIntToString string = strconv.Itoa(5000)
	fmt.Println(resultIntToString) // => "5000"
}
