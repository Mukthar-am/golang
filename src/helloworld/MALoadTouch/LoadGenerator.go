package matouchevents

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"time"
	"sync/atomic"
	"helloworld/masync"
	//"github.com/aerospike/aerospike-client-go/types/atomic"
)


var Cntr masyncatomic.AtomicInt

var Counter uint64 = 0

func GetCounter() (int)  {
	return Cntr.Get()
}

func GetMyCounter() (uint64) {
	return atomic.LoadUint64(&Counter)
}

func ResetAndPost() {
	// Reset event counter
	resetUrl := "http://localhost:8080/springwebeg/track/reset"
	execGetUri(resetUrl)

	url := "http://localhost:8080/springwebeg/track/info"
	_ = execGetUri(url)
	//log.Println("# Events := ", cnt)


	for {
		time.Sleep(time.Second * 1)
		//cntr.AddAndGet(1)
		atomic.AddUint64( &Counter, 1 )

		Cntr.GetAndIncrement()


		maTrackUrl := "http://localhost:8080/springwebeg/track/track"
		PostEvents(maTrackUrl)

		//fmt.Println("# Counter = ", Counter, ", Cntr = ", Cntr)
	}
}

func PostEvents(maTrackUrl string)  {
	m := map[string]interface{}{
		"name": "backy",
		"age": "33",
		"city": "Bangalore",
		"description": "GoLang here!",
	}
	mJson, _ := json.Marshal(m)
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
