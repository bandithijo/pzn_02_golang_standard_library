package main

import (
	"bufio"
	"os"
)

/*
Package bufio
- Package bufio atau singkatan dari buffered io
- Package ini digunakan untuk membuat data IO seperti Reader dan Writer
- <https://pkg.go.dev/bufio>
*/

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Hello, World!\n")
	_, _ = writer.WriteString("Selamat Belajar Go\n")

	writer.Flush()
	// Hello, World!
	// Selamat Belajar Go
}
