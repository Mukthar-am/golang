package main

import "fmt"

func main() {
	// if - Condition
	i := 1
	if i == 1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}


	// Loop'd for
	for i <= 10 {
		fmt.Print( "\ni=", i )
		i++
	}

	/** Conditional for */
	for i := 1; i <= 10; i++ {
		fmt.Print("\ni = ", i)
	}
}
