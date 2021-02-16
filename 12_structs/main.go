package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
	// or firstname lastname string
	//age intr
}

//function pass by value,greeting method(value reciever)

func (p Person) greet() string {
	return "hello , my name is " + p.firstname + " " + p.lastname + " and I am " + strconv.Itoa(p.age)
} //p -> identifier

//hasBirthday method  (pointer reciever)
func (p *Person) hasBirthday() {
	p.age = 5
}

func main() {
	//init person using struct

	person1 := Person{firstname: "Parth", lastname: "arora", city: "bhopal", gender: "m", age: 25}
	fmt.Println(person1)
	//alternative
	person2 := Person{"Parth", "arora", "bhopal", "m", 25}
	person1.age = 21
	fmt.Println(person1, person2, person1.firstname)

	person1.hasBirthday()
	fmt.Println(person1.greet())

}
