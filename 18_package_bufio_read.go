package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
Package bufio
- Package bufio atau singkatan dari buffered io
- Package ini digunakan untuk membuat data IO seperti Reader dan Writer
- <https://pkg.go.dev/bufio>
*/

func main() {
	input := strings.NewReader("This is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
		// This is long string
		// with new line
	}
}
