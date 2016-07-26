package loadgen

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"time"
)


func PostEventsCatchStatus(maTrackUrl string)  {
	m := map[string]interface{}{
		"name": "backy",
		"age": "33",
		"city": "Bangalore",
		"description": "GoLang here!",
	}
	mJson, _ := json.Marshal(m)
	s := string(mJson)

	request := gorequest.New().Timeout(1000 * time.Millisecond)
	resp, body, errs := request.Post(maTrackUrl).
	Set("Content-Type", "application/json").
	Send(s).
	End()

	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Print("\n# HttpStatus := ",resp.Status,"Body := ", body)
}

/**	==========================================================================================
	Using http.post client
 */
func PostEventsByHttpClient(URL string) {
	m := map[string]interface{}{
		"name": "Muks",
		"age": "33",
		"city": "Bangalore",
	}
	mJson, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJson)
	resp, err := http.Post(URL, "application/json", contentReader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print out
	fmt.Println("Event Post Response := ", string(htmlData))
}
