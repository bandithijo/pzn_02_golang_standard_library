package main

/*
Package io
- IO atau singkatan dari Input Output, merupakan fitur di Go yang digunakan sebagai standard untuk proses Input Output
- Di Go, semua mekanisme input output pasti mengikuti standard Package io
- <https://pkg.go.dev/io>

Reader
- Untuk membaca input, Go menggunakan kontrak interface bernama Reader yang terdapat di package io

type Reader interface {
	Read(p []byte) (n int, err error)
}

Writer
- Untuk menulis ke output, Go menggunakan kontrak interface bernama Writer yang terdapat di package io

type Writer interface {
	Write(p []byte) (n int, err error)
}

Implementasi IO
- Penggunaan dari IO sendiri di Go, terdapat di banyak package, sebelumnya contohnya kita menggunakan CSV Reader dan CSV Writer
- Karena Package IO sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri
- Tapi untuknya, Go juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu menggunakan package bufio
*/

