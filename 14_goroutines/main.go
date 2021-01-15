package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { //1st go routine
	var wg sync.WaitGroup
	wg.Add(1) //increment by 1

	go func() {
		count("sheep")

		wg.Done() //decrement
	}()
	wg.Wait() //wait till 0
	//go count("sheep") //go run next step and run this in the background

}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
