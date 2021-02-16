package main

import (
	"fmt"
)

func main() {
	//Arrays

	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	//declear and assign

	fuit := [2]string{"apple", "orange"}

	fmt.Println(fruitArr[1], fuit)

	//slice

	fruitSlice := []string{"Apple", "orange"}

	///new declearation   fruitslice := make([]int,5,10)  //5 values 0 and size =10  and also fruit : =append(fruit , 0 ,-1) for the remain values 

	fmt.Println(fruitSlice, len(fruitSlice), fruitSlice[1:2])
	//count length of slice

	//1::2 starts at one and stops at 2

}
