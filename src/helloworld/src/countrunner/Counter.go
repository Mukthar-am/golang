package main

import (
	"log"
	"time"
	"matouchload/eventsloader"
	//"github.com/aerospike/aerospike-client-go/types/atomic"
	//"helloworld/masync"
	"fmt"
)

func main() {
	//log.Println("# EventsCount := ", MATouchload.Ops)
	//log.Println("# MATouchload.Ops := ", MATouchload.Ops)
	////c := make(chan int)
	//GoCounterRoutine(c)


	//fmt.Println(MATouchLoad.GetCounter())
	//
	//MATouchLoad.PrintCounter()
	//type cntr masyncatomic.AtomicInt
	//fmt.Println(matouchevents.GetCounter())
	fmt.Println(matouchevents.GetEventsCount())

}



func GoCounterRoutine(ch chan int) {
	counter := 0
	for {
		ch <- counter
		counter += 1

		log.Println(counter)
	}
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
