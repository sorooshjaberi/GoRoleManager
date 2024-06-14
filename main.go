package main

import (
	"fmt"
	"sync"
	"time"
)

var names = []string{"Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran"}

func main() {
	dataChan := make(chan string)
	signalCreateReader := make(chan bool)

	wg := new(sync.WaitGroup)

	writer := func(wg *sync.WaitGroup, dataChan chan<- string) {

		defer wg.Done()
		defer close(dataChan)
		defer close(signalCreateReader)

		// wait for 1 second and if delay,
		// aka. tolerance
		timeout := time.Millisecond * 500

		for _, name := range names {
			select {
			case dataChan <- name:
			case <-time.After(timeout):
				// if timed out
				signalCreateReader <- true
			}
		}
	}
	reader := func(wg *sync.WaitGroup, dataChan <-chan string, timeout time.Duration) {
		defer wg.Done()
		defer fmt.Println("Bye!!!!!!")
		 fmt.Println("Hey !!!!!!")
		for {
			// reader delay
			time.Sleep(time.Second)

			select {
			case name, ok := <-dataChan:
				if !ok {
					return
				}
				fmt.Printf("name: %v\n", name)
			// timeout for idle reader
			case <-time.After(timeout):
				return
			}
		}

	}
	wg.Add(1)
	go writer(wg, dataChan)

	for range signalCreateReader {
		wg.Add(1)
		go reader(wg, dataChan, time.Millisecond*1000)
	}

	wg.Wait()

}
