package main

import (
	"fmt"
	"time"
)

func main() { //1st go routine

	c1 := make(chan string) //channel
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "e"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select { //keep taking channels which are available
		case msg1 := <-c1:
			fmt.Println((msg1))

		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
