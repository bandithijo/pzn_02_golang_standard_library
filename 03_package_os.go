package main

import (
	"fmt"
	"os"
)

/*
Package os
- Go telah menyediakan banyak sekali pacakge bawaan, salah satunya adalah package os
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi independen (bisa digunakan disemua sistem operasi)
- <https://pkg.go.dev/os>
*/

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
