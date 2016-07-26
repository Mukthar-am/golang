package main

import (
	"matouchload/configs"
	"github.com/creamdog/gonfig"
	"fmt"
	"encoding/json"
)

type file gonfig.Gonfig

func main() {
	cfgAbsPath := "/Users/15692/Data/git/golang/src/matouchload/configs/TouchloadConfigs.gonfig"

	file := confighelper.GetFileHandler(cfgAbsPath)
	uname, _ := file.Get("payload/test", "empty-json")
	//payload := map[string]interface{}{
	//	"name": "backy",
	//	"age": "33",
	//	"city": "Bangalore",
	//	"description": "GoLang here!",
	//}
	mJson, _ := json.Marshal(uname)
	s := string(mJson)

	fmt.Println("======",uname)

	//mJson, _ := json.Marshal(uname)
	//s := string(mJson)

	fmt.Println("# Load:= ", s)

	url, _ := file.Get("url/test", "empty-json")
	fmt.Println("======",url)
}
