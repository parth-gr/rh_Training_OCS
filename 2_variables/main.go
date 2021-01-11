package main

import "fmt"

var global__var  = 32   //global 

func main(){                    //opening should be in this line
	var name1 string = "Parth" //decleration1
	var name2  int32 = 32      //decleration2
	const name3 = "parth"   //not redifined
    name4 := "arora"   //equivalent to ' var name1 = "arora" '  
	size := 1.3        //float
	

	name, surname := "parth" , "arora"  //one line declearation

	fmt.Println(name1, name2 ,name3 ,global__var , name4 ,name,surname) // any variable should be used compolsury
    fmt.Printf("%T\n" ,size)  //fmt operator %T => type of
}