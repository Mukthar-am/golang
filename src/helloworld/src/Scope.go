package main

import "fmt"

var foo string = "global"
var Counter int = 3

func main() {
	fmt.Println(foo) // prints "global"
	fmt.Println("#1 =>",Counter)

	// using := creates a new function scope variable
	// named foo that shadows the package scope foo
	foo := "function scope"
	Counter = Counter + 1
	fmt.Println(foo) // prints "function scope"
	fmt.Println("#1 => ", Counter)

	printGlobalFoo() // prints "global"

	if true {
		foo := "nested scope"
		fmt.Println(foo) // prints "nested scope"

		printGlobalFoo() // prints "global"
	}
	// the foo created inside the if goes out of scope when
	// the code block is exited

	fmt.Println(foo) // prints "function scope"
	printGlobalFoo() // prints "global"

	if true {
		foo = "nested scope" // note just = not :=
	}

	fmt.Println(foo) // prints "nested scope"
	printGlobalFoo() // prints "global"

	setGlobalFoo()
	printGlobalFoo() // prints "new value"
}

func printGlobalFoo() {
	fmt.Println("")
	fmt.Println(foo)
	fmt.Println("#3 => ", Counter)
}

func setGlobalFoo() {
	foo = "new value" // note just = not :=
	Counter = 1111
}