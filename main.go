package main

import (
	"fmt"
	"github.com/cagox/fluxspellsapi/crypto"
)

func main() {
	fmt.Println("Hello World!")

	string1 := crypto.HashPassword("password1")
	string2 := crypto.HashPassword("password2")
	string3 := crypto.HashPassword("password3")
	string4 := crypto.HashPassword("password1wasnotlognenough")
	string5 := crypto.HashPassword("Mary Had a Little Lamb")

	fmt.Println(len(string1))
	fmt.Println(len(string2))
	fmt.Println(len(string3))
	fmt.Println(len(string4))
	fmt.Println(len(string5))

}
