package main

import (
	"fmt"
	"time"
)

/*
Package time
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go
- <https://pkg.go.dev/time>

Beberapa Function di Package Time
- time.Now()
  Untuk mendapatkan waktu saat ini
- time.Date(...)
  Untuk membuat waktu
- time.Parse(layout, string)
  Untuk memparsing waktu dari string

Duration
- Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
- Package time memiliki type Duration, yang sebenarnya adalah alias untuk int64
- Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration
*/

func main() {
	// time
	var now time.Time = time.Now()
	fmt.Println(now.Local()) // => 2025-01-15 01:27:56.367404605 +0800 WITA

	var utc time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local()) // => 2009-11-11 07:00:00 +0800 WITA

	parse, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err == nil {
		fmt.Println(parse) // => 2006-01-02 15:04:05 +0000 UTC
	} else {
		fmt.Println("Error:", err.Error())
	}

	// duration
	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2

	fmt.Printf("duration1: %d\n", duration1) // => duration1: 100000000000
	fmt.Printf("duration2: %d\n", duration2) // => duration2: 10000000
	fmt.Printf("duration3: %d\n", duration3) // => duration3: 99990000000

	fmt.Println("duration1:", duration1) // => duration1: 1m40s
	fmt.Println("duration2:", duration2) // => duration2: 10ms
	fmt.Println("duration3:", duration3) // => duration3: 1m39.99s
}
