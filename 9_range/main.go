package main

import "fmt"

func main() {
	//loop through array , slice , map

	arr := []int{1, 2, 3, 4, 5, 6}

	for i, id := range arr {
		fmt.Printf("%d -ID : %d\n", i, id)
	}

	//not using index

	for _, id := range arr { //using _ if i not used
		fmt.Printf("ID : %d\n", id)
	}
	var sum = 0
	for _, id := range arr {
		sum += id
	}
	fmt.Println(sum)

	//Range with map

	name := map[string]string{"parth": "par", "sharon": "sh"}

	for k, v := range name {
		fmt.Printf("%s %s ", k, v)
	}

	for k := range name {
		fmt.Printf("%s  ", k)
	}
}
