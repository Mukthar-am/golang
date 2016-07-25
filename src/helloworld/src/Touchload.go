package main

import (
	"log"
	"helloworld/maloadtouch"
	"fmt"
)

func main() {
	go sendLoad()
	var input string
	fmt.Scanln(&input)
	fmt.Println("Terminate request received: ", input)
}

func sendLoad() {
	url := "http://localhost:8080/springwebeg/track/info"
	cnt := MATouchload.EventsCounter(url)
	log.Println("# Events := ", cnt)

	for {
		maTrackUrl := "http://localhost:8080/springwebeg/track/events"
		MATouchload.SendEvents(maTrackUrl)
	}
}
