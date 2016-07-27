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

/* CLI args */
var concUsers int
var logFile string
var configFile string

var testmode bool // to run in test mode


func init() {
	logPathDefault, _ := filepath.Abs("touchload.log")
	configDefault, _ := filepath.Abs("../../configs/TouchloadConfigs.gonfig")

	flag.IntVar(&concUsers, "runs", 1, "help message for flagname")
	flag.StringVar(&logFile, "log", logPathDefault, "log file path")
	flag.StringVar(&configFile, "config", configDefault, "config file path")
	flag.BoolVar(&testmode, "testmode", true, "run in test mode")

	flag.Parse()
}

/**
	golang - MaTouchload launcher script
 */
func main() {
	fmt.Print("\n")
	log.Print(":= (START TIME) =:")
	fmt.Println(":= Load Start :=",
		"\nUsers/parallel runs -> ", concUsers,
		"\nLog file: ", logFile,
		"\nConfig file: ", configFile)

	file := configs.GetFileHandler(configFile)
	url, _ := file.Get("url/test", "empty-json")
	maTrackingUrl = url.(string)

	load, _ := file.Get("payload-track", "empty-json")

	// concurrently hitting the url
	go sendLoad(maTrackingUrl, load)

	select {}
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println("Terminate request received: ", input)
}

func sendLoad(maTrackingUrl string, payload interface{}) {
	// open a file
	logFH, err := os.OpenFile(logFile, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer logFH.Close()
	log.SetOutput(logFH)        // assign it to the standard logger

	fmt.Println("===========================================================")
	fmt.Println("# Runtime stats:= ")
	for i := 1; i <= concUsers; i++ {
		go touchloader.ResetPoster(maTrackingUrl, testmode, payload)
	}

	type Counters2xx int
	type Counters5xx int
	for {
		Counters2xx := touchloader.GetCount2xx()
		Counters5xx := 0

		time.Sleep(time.Second * 2)
		fmt.Printf("\r(2xx/5xx) := %d/%d", Counters2xx, Counters5xx)
		log.Println(Counters2xx)
	}
}
