package main

import "fmt"

func main() {

	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y) //%d base 10 ints
	} else {
		fmt.Printf("hello\n")
	}

	//Switch
	color := "red"

	switch color {
	case "red":
		fmt.Println("red")
	case "blue":
		fmt.Println("blue")
	default:
		fmt.Println("no")
	}

}
