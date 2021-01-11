package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int { //or (num1 , num2 int) both are same
	return num1 + num2
}
func main() {

	fmt.Println(greeting("Parth"), getSum(2, 3))
}
