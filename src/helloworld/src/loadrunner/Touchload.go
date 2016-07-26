package main

import (
	"helloworld/maloadtouch"
	"fmt"
	"time"
)

func main() {

	go sendLoad()
	var input string
	fmt.Scanln(&input)
	fmt.Println("Terminate request received: ", input)
}


func sendLoad() {
	var Counters5xx int = 0
	fmt.Println("===========================================================")
	fmt.Println("# Runtime stats:= ")
	go matouchevents.ResetAndPost()
	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("\r(2xx/5xx) := %d/%d", matouchevents.GetMyCounter(), Counters5xx)
	}

}

