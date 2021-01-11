package main

import ( //multiple packages
	"fmt"
	"math"

	"github.com/parth-gr/rh_Training_OCS/3_packages/usingpackage" //path after the source
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	var num float64 = 4
	ans := usingpackage.Sqrtofno(num)
	fmt.Println(ans)

}
