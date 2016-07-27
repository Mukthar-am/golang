package main

import (
	"fmt"
	"time"
	"github.com/creamdog/gonfig"
	"log"
	"flag"
	"matouchload/touchlibs/configs"
	"matouchload/touchlibs/touchloader"
	"os"
	"path/filepath"
)

type file gonfig.Gonfig
var maTrackingUrl string
var runs int
var logFile string
var configFile string


func init() {
	logPathDefault, _ := filepath.Abs("touchload.log")
	configDefault, _ := filepath.Abs("../../configs/TouchloadConfigs.gonfig")

	flag.IntVar(&runs, "runs", 1, "help message for flagname")
	flag.StringVar(&logFile, "log", logPathDefault, "log file path")
	flag.StringVar(&configFile, "config", configDefault, "config file path")
	flag.Parse()
}

func main() {


	fmt.Print("\n")
	log.Print(":= (START TIME) =:")
	fmt.Println(":= Load Start :=",
	"\nUsers/parallel runs -> ", runs,
	"\nLog file: ", logFile,
	"\nConfig file: ", configFile)

	os.Exit(0)

	file := configs.GetFileHandler(configFile)
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

	// open a file
	f, err := os.OpenFile("/Users/15692/Downloads/test.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)


	var testmode bool = true
	var Counters5xx int = 0


	fmt.Println("===========================================================")
	fmt.Println("# Runtime stats:= ")

	go touchloader.ResetPoster(maTrackingUrl, testmode, payload)

	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("\r(2xx/5xx) := %d/%d", touchloader.GetEventsCount(), Counters5xx)
		log.Println("mukthar")
		//f.WriteString("blah asfsjk")

	}



}
