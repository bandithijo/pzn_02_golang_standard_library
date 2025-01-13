package main

import "fmt"

/*
Package fmt
- Sebelumnya kita sudah sering menggunakna package fmt dengan menggunakan function Println
- Selain Println, masih banyak function yang terdapat di package fmt, contohnya banyak digunakan untuk melakukan format
- <https://pkg.go.dev/fmt>
*/

func main() {
	fmt.Println("Hello, World!")

	firstName := "Rizqi"
	lastName := "Assyaufi"

	fmt.Printf("Hello, '%s %s'!\n", firstName, lastName)
}
