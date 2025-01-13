package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
Package container/ring
- Package container/ring adalah implementasi struktur data circular list
- Circular list adalah struktur data ring, dimana di akhir element akan kembali ke element awal (HEAD)
- <https://pkg.go.dev/container/ring>
*/

func main() {
	// membuat data ring
	var data *ring.Ring = ring.New(5)

	// mengisi data manual
	// data.Value = "Value 1"
	// data = data.Next()
	// data.Value = "Value 2"
	// data = data.Next()
	// data.Value = "Value 3"
	// data = data.Next()
	// data.Value = "Value 4"
	// data = data.Next()
	// data.Value = "Value 5"

	// mengisi data sejumlah panjang variable data
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// menampilkan isi data dengan method .Do() dari ring
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
