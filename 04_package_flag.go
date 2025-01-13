package main

import (
	"flag"
	"fmt"
)

/*
Package flag
- Package flag berisikan fungsionalitas	untuk memparsing command line argument
- <https://pkg.go.dev/flag>
*/

func main() {
	// flag, balikannya adalah pointer
	var host *string = flag.String("host", "localhost", "Put your database host")
	var port *string = flag.String("port", "0", "Put your database port")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")

	// untuk melakukan parsing
	flag.Parse()

	// maka gunakan asterisk operator ketika ingin diprint
	// $ go run 04_package_flag.go -host=123.123.123.1 -port=5432 -username=bandithijo -password=secret
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}
