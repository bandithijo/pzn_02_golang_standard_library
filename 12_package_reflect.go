package main

import (
	"fmt"
	"reflect"
)

/*
Package reflect
- Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
- Hal ini bisa dilakukan di Go dengan menggunakan package reflect
- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa bereksplorasi dengan package reflect ini secara otodidak
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
- <https://pkg.go.dev/reflect>
*/

type Sample struct {
	Code string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Age     int `required:"true" max:"100"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println("Field", structField.Name, "with Type", structField.Type)
		fmt.Println("required:", structField.Tag.Get("required"))
		fmt.Println("max:", structField.Tag.Get("max"))
	}
}

func IsValid(data any) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	readField(Sample{"Golang"})
	// Type Name Sample
	// Field Code with Type string
	// required: true
	// max: 10

	readField(Person{"Rizqi", "Balikpapan", 36})
	// Type Name Person
	// Field Name with Type string
	// required: true
	// max: 10
	// Field Address with Type string
	// required: true
	// max: 10
	// Field Age with Type int
	// required: true
	// max: 100

	person := Person{
		Name: "Rizqi",
		Address: "Balikpapan",
		Age: 35,
	}
	fmt.Println("is valid:", IsValid(person)) // => is valid true

	person = Person{
		Name: "Rizqi",
		Address: "",
		Age: 35,
	}
	fmt.Println("is valid:", IsValid(person)) // => is valid false
}
