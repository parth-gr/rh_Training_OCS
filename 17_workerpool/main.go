package main

import ("fmt" 
       "time"
	)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	c := make(chan int, 45)
	c1 := make(chan int, 45)
	go worker(c, c1)
	go worker(c, c1)
	go worker(c, c1)
	go worker(c, c1)

	

	for i := 0; i < 45; i++ {
		c <- i
	}
	close(c)

	for i := 0; i < 45; i++ {
		fmt.Println(i,<-c1)
	}
	currentTime1 := time.Now()
	fmt.Println(currentTime1.Format("15:04:05"))
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
