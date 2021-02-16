package main

import (
	"fmt"
	"strings"
	//      "sort"
	//       "os"
	//      "log"
	//      "io/ioutil"
	//     "strconv"
)

func main() {
	fmt.Println("hello world")

	a := "hell"
	fmt.Println(strings.Contains(a, "he"))
     fmt.Println(strings.Index(a, "he"))
     fmt.Println(strings.Count(a, "he"))
     fmt.Println(strings.Replace(a, "e" , "p" , 1)) //first 1's

}
