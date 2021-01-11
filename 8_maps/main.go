package main

import "fmt"

func main() {
	//define map

	emails := make(map[string]string) //first datatype for key and then for value

	//Assign key values

	emails["parth"] = "partharora@gmail.com"
	emails["par"] = "par@gmail.com"
	emails["r"] = "r@gmail.com"

	fmt.Println(len(emails)) //length of map
	fmt.Println(emails["par"])

	delete(emails, "r") //delete a key

	fmt.Println(emails)

	//declear map and intialize

	name := map[string]string{"parth": "par", "sharon": "sh"}
	name["r"] = "r@gmail.com"
	fmt.Println(name)

}
