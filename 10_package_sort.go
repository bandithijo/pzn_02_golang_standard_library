package main

import (
	"fmt"
	"sort"
)

/*
Package sort
- Package sort adalah package yang berisikan utilitas untuk proses pengurutan
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
- <https://pkg.go.dev/sort>

// elements of the collection be enumerated by an integer index
type Interface interface {
	// Len is the number of elements in the collection
	Len() int
	// Less reports whether the element with index i should sort before the
	// element with index j
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j
	Swap(i, j int)
}
*/

type User struct {
	Name string
	Age int
}

type UserSlice []User // alias type declaration untuk slice User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Rizqi", 26},
		{"Nur", 18},
		{"Assyaufi", 32},
	}

	fmt.Println(users) // => [{Rizqi 26} {Nur 18} {Assyaufi 32}]

	sort.Sort(UserSlice(users))

	fmt.Println(users) // => [{Nur 18} {Rizqi 26} {Assyaufi 32}]
}
