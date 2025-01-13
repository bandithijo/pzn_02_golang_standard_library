package main

import (
	"fmt"
	"strings"
)

/*
Package strings
- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
- Ada banyak sekali function yang bisa kita gunakan
- <https://pkg.go.dev/strings>

Beberapa Function di Package strings
- strings.Trim(string, cutset)
  Memotong cutset di awal dan akhir string
- strings.ToLowe(string)
  Membuat semua karakter string menjadi lower case
- strings.ToUpper(string)
  Membuat semua karakter string menjadi upper case
- strings.Split(string, separator)
  Memotong string berdasarkan separator
- strings.Contains(string, substring)
  Mengecek apakah string mengandng string lain
- strings.ReplaceAll(string, from, to)
  Mengubah semua string dari from ke to
*/

func main() {
	fmt.Println(strings.Contains("Rizqi Nur Assyaufi", "nur")) // => false
	fmt.Println(strings.Contains("Rizqi Nur Assyaufi", "Nur")) // => true
	fmt.Println(strings.ContainsAny("Rizqi Nur Assyaufi", "nur")) // => true
	fmt.Println(strings.Contains("Rizqi Nur Assyaufi", "")) // => true
	fmt.Println(strings.ContainsAny("Rizqi Nur Assyaufi", "")) // => false

	fmt.Println(strings.Split("Rizqi Nur Assyaufi", "|")) // => [Rizqi Nur Assyaufi]

	fmt.Println(strings.ToLower("Rizqi Nur Assyaufi")) // => rizqi nur assyaufi
	fmt.Println(strings.ToUpper("Rizqi Nur Assyaufi")) // => RIZQI NUR ASSYAUFI

	fmt.Println(strings.Trim("   Rizqi Nur Assyaufi   ", " ")) // => Rizqi Nur Assyaufi

	fmt.Println(strings.ReplaceAll("Rizqi Nur Rizqi Assyaufi", "Rizqi", "R.")) // => R. Nur R. Assyaufi
}
