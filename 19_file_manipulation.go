package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
File Manipulation
- Di Package os, terdapat File Management
- Saat kita membuat atau membaca file menggunakan Package os, struct File merupakan implementasi dari io.Reader dan io.Writer
- Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package io/bufio

Open File
- Untuk membuat/membaca File, kita bisa menggunakan os.OpenFile(name, flag, permission)
- name berisikan nama file, bisa absolute atau relative/local
- flag merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain
- permission merupakan permission yang diperlukan ketika membuat file, bisa kita simulasikan di sini: <https://chmod-calculator.com>

const (
	// Exactly one of O_RDONLY, OWRONLY, O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.

	// The remaining values may be or'ed in the control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing
	O_CREATE int = syscall.O_CREATE // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
*/

func createNewFile(name string, message string) error {
	// create new file
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// write into file
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	// open file
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read file
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	// open file
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// write into file
	file.WriteString(message)
	return nil
}

func main() {
	var targetFile string = "sample.log"
	var messages string = "this is sample log.\nhow to create and write file in go.\n"

	// create file and write
	createNewFile(targetFile, messages)

	// append into file
	var appendMessage string = "this is appended text."
	addToFile(targetFile, appendMessage)

	// read file
	result, err := readFile(targetFile)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}
