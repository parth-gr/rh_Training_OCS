package main

import (
	"fmt"
	"time"
)

func main() { //1st go routine

	c := make(chan string) //channel
	go count("sheep", c)   //go run next step and run this in the background

	for msg := range c {
		//msg := <-c //recieve the message

		fmt.Println(msg)
	}

	//sending and recieving are blocking operation ,

	//buffered channel

	c1 := make(chan string, 2)
	c1 <- "hello"
	msg := <-c1
	fmt.Println(msg)
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		//fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c) //sender must only close it
}
