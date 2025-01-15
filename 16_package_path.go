package main

import (
	"fmt"
	"path"
)

/*
Package path
- Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File System
- Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok untuk data URL
- <https://pkg.go.dev/path>
- Namun jika ingin menggunakan untuk memanipulasi path di File System, karena Windows menggunakan backslack, maka khusus untuk File System, perlu menggunakan package path/filepath
- <https://pkg.go.dev/path/filepath>
*/

func main() {
	// path
	fmt.Println(path.Dir("hello/world.go")) // => hello
	fmt.Println(path.Base("hello/world.go")) // => world.go
	fmt.Println(path.Ext("hello/world.go")) // => .go
	fmt.Println(path.Join("hello", "world", "main.go")) // => hello/world/main.go
}
