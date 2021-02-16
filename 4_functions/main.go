package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int { //or (num1 , num2 int) both are same
	return num1 + num2
}
func main() {
	// defer func1() defer function (only 1)will be called at the lat of the main function overs
	 //  func2()   

	 //defer func(){
		// fmt.Println(recover())
	 //}()
	fmt.Println(greeting("Parth"), getSum(2, 3))
}


