package touchloader

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"time"
	"sync/atomic"
)

var Counter uint64 = 0


/* ===========================================================================================
	Get the total number of concurrent events count
	*/
func GetEventsCount() (uint64) {
	return atomic.LoadUint64(&Counter)
}

/* ===========================================================================================
	# Totally for testing purpose
	- Reset event counts and
	- Post events
	*/
func ResetPoster(maTrackingUrl string, testmode bool, payload interface{}) {
	if (testmode == true) {
		// Reset event counter
		resetUrl := "http://localhost:8080/springwebeg/track/reset"
		execGetUri(resetUrl)

		url := "http://localhost:8080/springwebeg/track/info"
		_ = execGetUri(url)
	}


	for {
		time.Sleep(time.Second * 1)
		atomic.AddUint64( &Counter, 1 )
		Poster(maTrackingUrl, payload)
	}
}

func PostOnly(postUrl string, payload string) {
	for {
		//time.Sleep(time.Second * 1)
		atomic.AddUint64( &Counter, 1 )
		Poster(postUrl, payload)
	}
}

func Poster(maTrackUrl string, payload interface{})  {
	//payload := map[string]interface{}{
	//	"name": "backy",
	//	"age": "33",
	//	"city": "Bangalore",
	//	"description": "GoLang here!",
	//}
	mJson, _ := json.Marshal(payload)
	s := string(mJson)

	request := gorequest.New().Timeout(1000 * time.Millisecond)
	_, _, errs := request.Post(maTrackUrl).
	Set("Content-Type", "application/json").
	Send(s).
	End()

	if errs != nil {
		fmt.Println(errs)
	}
}



func execGetUri(URL string) (string) {
	resp, err := http.Get(URL)
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
	//fmt.Println("EventsCount:=", string(htmlData))
	return string(htmlData)
}

