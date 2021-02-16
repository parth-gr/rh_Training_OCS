package main

import "fmt"

func main() {
	//location address of a variable

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) //*int

	fmt.Println(*b)

	fmt.Println(*&a)

	//change val by pointer

	*b = 6

	fmt.Println(a)
	//passing ptrs while big data is there then performance is increased
	//only pass by value exists

	// p := new(int)
	// *p =100

}
