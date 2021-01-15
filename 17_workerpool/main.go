package main

import "fmt"

func main() {

	c := make(chan int, 100)
	c1 := make(chan int, 100)
	go worker(c, c1)
	go worker(c, c1)
	go worker(c, c1)
	go worker(c, c1)

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)

	for i := 0; i < 100; i++ {
		fmt.Println(<-c1)
	}
}

func worker(c <-chan int, c1 chan<- int) {
	for n := range c {
		c1 <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return (fib(n-2) + fib(n-1))
}
