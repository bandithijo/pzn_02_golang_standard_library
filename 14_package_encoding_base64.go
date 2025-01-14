package main

import (
	"encoding/base64"
	"fmt"
)

/*
Package encoding
- Go menyediakan package encoding untuk melakukan encode dan decode
- <https://pkg.go.dev/encoding>
- Go menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah base64, csv, dan json
*/

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Rizqi Nur Assyaufi"))
	fmt.Println(encoded) // => Uml6cWkgTnVyIEFzc3lhdWZp

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Println(decoded) // => [82 105 122 113 105 32 78 117 114 32 65 115 115 121 97 117 102 105]
		fmt.Println(string(decoded)) // => Rizqi Nur Assyaufi
	} else {
		fmt.Println("Error:", err.Error())
	}
}
