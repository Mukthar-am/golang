package main

import (
	"fmt"
	"time"
	"math/rand"
)

func concFunc(n int) {
	for i := 0; i < 10; i++ {
		time.Sleep( time.Duration(rand.Intn(1000)) )
		fmt.Print("\n", time.Now() ," := Process ID # ", n, ", Value = ", i)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		go concFunc(i)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("Terminate request received: ", input)
}
