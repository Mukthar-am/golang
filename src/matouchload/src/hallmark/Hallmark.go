package main

import (
	"matouchload/eventsloader"
	"fmt"
	"time"
	"matouchload/configs"
	"github.com/creamdog/gonfig"
)

type file gonfig.Gonfig
var cfgAbsPath = "/Users/15692/Data/git/golang/src/matouchload/configs/TouchloadConfigs.gonfig"

var maTrackingUrl string
var payload string

func main() {
	file := confighelper.GetFileHandler(cfgAbsPath)
	url, _ := file.Get("url/test", "empty-json")
	maTrackingUrl = url.(string)

	load, _ := file.Get("payload/test", "empty-json")

	// concurrently hitting the url
	go sendLoad(maTrackingUrl, load)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Terminate request received: ", input)
}

func sendLoad(maTrackingUrl string, payload interface{}) {
	var testmode bool = true
	var Counters5xx int = 0


	fmt.Println("===========================================================")
	fmt.Println("# Runtime stats:= ")

	go touchloader.ResetPoster(maTrackingUrl, testmode, payload)

	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("\r(2xx/5xx) := %d/%d", touchloader.GetEventsCount(), Counters5xx)
	}
}
