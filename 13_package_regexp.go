package main

import (
	"fmt"
	"regexp"
)

/*
Package regexp
- Package regexp adalah utilitas di Go untuk melakukan pencarian regular expression
- Regular expression di Go menggunakan library C yang dibuat Google bernama RE2
- <https://github.com/google/re2/wiki/Syntax>
- <https://pkg.go.dev/regexp>

Beberapa Function di Package regexp
- regexp.MustCompoile(string)
  Membuat Regexp
- Regexp.MatchString(string) bool
  Mengecek apakah Regexp match dengan string
- Regexp.FindAllString(string, max)
  Mencari string yang match dengan maximum jumlah hasil
*/

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`r([a-z])i`)

	fmt.Println(regex.MatchString("rizqi")) // => false
	fmt.Println(regex.MatchString("rzi")) // => true
	fmt.Println(regex.MatchString("rZi")) // => false

	fmt.Println(regex.FindAllString("rizqi Rizqi r1i rei r3i rzq rzi rZq rii", 10)) // => [rei rzi rii]
}
