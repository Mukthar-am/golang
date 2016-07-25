package main

import (
	"fmt"
	"helloworld/lib"
)

func main() {
	fmt.Println(lib.Reverse("!oG ,olleH\n\n"))

	var name string = "Mukthar" /* String */
	fmt.Print("My name is ", name)

	var age int = 33	/* Integer */
	fmt.Print(" and his age is ", age)

	/** Defining multiple variables */
	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Print("\na=",a,", b=", b, ", c=", c)
}
