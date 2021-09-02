package main

import (
	"fmt"
	"sort"
)

func main() {
	sli1 := make([]int, 3)
	fmt.Println("Please enter integer; enter X to exit : ")
	fmt.Println("slice length is : ")
	fmt.Println(len(sli1))
	fmt.Println("current array is : ")
	for _, val := range sli1 {
		fmt.Println(val)
	}
	fmt.Println("-------------------")
	for 1 > 0 {

		var i int
		_, err := fmt.Scanf("%d", &i)

		if err != nil {
			fmt.Println("Exiting program as X/x or non-int is entered!  ")
			break
		}
		// fmt.Println("u input : ")
		// fmt.Println(i)
		// for _, v := range i {
		// 	sli1 = append(sli1, v)
		// }
		sli1 = append(sli1, i)
		sort.Ints(sli1)
		fmt.Println("current sorted array is : ")
		fmt.Println(sli1)
		fmt.Println("-------------------")
	}
}
