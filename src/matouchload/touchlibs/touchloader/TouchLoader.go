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

var Counter2xx uint64 = 0
var CounterNon2xx uint64 = 0


/* ===========================================================================================
	Get the total number of concurrent events count
	*/
func GetCount2xx() (uint64) {
	return atomic.LoadUint64(&Counter2xx)
}

func GetCount5xx() (uint64) {
	return atomic.LoadUint64(&CounterNon2xx)
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
		time.Sleep(time.Second * 1)	// comment this in realtime
		Poster(maTrackingUrl, payload)
	}
}

func PostOnly(postUrl string, payload string) {
	for {
		//time.Sleep(time.Second * 1)
		atomic.AddUint64( &Counter2xx, 1 )
		Poster(postUrl, payload)
	}
}

func Poster(maTrackUrl string, payload interface{})  {
	mJson, _ := json.Marshal(payload)
	s := string(mJson)

	request := gorequest.New().Timeout(1000 * time.Millisecond)
	resp, _, errs := request.Post(maTrackUrl).
	Set("Content-Type", "application/json").
	Send(s).
	End()

	if errs != nil {
		fmt.Println(errs)
	}

	statusCode := resp.StatusCode
	if (statusCode == 200) {
		atomic.AddUint64( &Counter2xx, 1 )
	} else {
		atomic.AddUint64( &CounterNon2xx, 1 )
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

