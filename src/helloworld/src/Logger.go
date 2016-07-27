package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")

	// open a file
	f, err := os.OpenFile("/Users/15692/Downloads/test.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	f.WriteString("blah asfsjk")

	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)
}