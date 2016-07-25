package maloadtouch

import (
	//"encoding/json"
	//"github.com/jmcvetta/napping"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
	"os"
	"time"
	"encoding/json"
)

func main() {

	//fmt.Print("Total Events := ")
	//getByClient()
	//url := "http://localhost:8080/springwebeg/track/info"
	//eventsCounter(url)
	//postByClient()
	//hitByGoReq()

	//urlByGoRequestApiTesting()
	//urlByGoRequestApiTracking()        /** url post request by gorequest()  */
}

func hitByGoReq() {
	contentType := "application/json"
	trackingUrl := "http://localhost:8080/springwebeg/track/events"
	fmt.Println("Tracking URL:>", trackingUrl)

	request := gorequest.New().Timeout(1000 * time.Millisecond)

	resp, body, errs := request.Post(trackingUrl).
	Set("Content-Type", contentType).
	Send(`{"name":"backy", "age":"dog", "city":"dog"}`).
	End()

	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", body)

}


func getByClient() {
	url := "http://localhost:8080/springwebeg/track/events"
	log.Println("(URL) :> ", url)

	client := &http.Client{}        // create client instance

	var jsonStr = []byte(`{"title":"dummy."}`)

	// create http request object
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

}

func postByClient() {
	m := map[string]interface{}{
		"name": "backy",
		"age": "33",
		"city": "Bangalore",
	}
	mJson, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", "http://localhost:8080/springwebeg/track/events", contentReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notes", "GoRequest is coming!")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

