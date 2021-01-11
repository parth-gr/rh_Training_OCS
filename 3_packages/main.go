package main

import (            //multiple packages
	"fmt" 
	"math"
	"3_packages/using_package"
)  

func main(){
	 fmt.Println(math.Floor(2.7))
	 fmt.Println(math.Ceil(2.7))
	 fmt.Println(math.Sqrt(2.7))
	 var num float64 = 4
	 ans := using_package.Sqrt_of_no(num)
	 fmt.Println(ans)
	 	
} 

