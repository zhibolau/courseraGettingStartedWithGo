/*

Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
 The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

 https://www.coursera.org/learn/golang-getting-started/peer/sLPZg/module-3-activity-slice-go/discussions/threads/I7lbsfCXEeuQQhJME_HymQ

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli1 := make([]int, 3)
	fmt.Println("Please enter integer; enter X to exit : ")
	// fmt.Println("slice length is : ")
	// fmt.Println(len(sli1))
	// fmt.Println("current array is : ")
	// for _, val := range sli1 {
	// 	fmt.Println(val)
	// }
	fmt.Println("-------------------")
	for 1 > 0 {

		// var i []int
		// _, err := fmt.Scanf("%d", &i)
		// if err != nil {
		// 	fmt.Println("Exiting program as X/x or non-int is entered!  ")
		// 	break
		// }
		targetStr := bufio.NewReader(os.Stdin)
		input, _ := targetStr.ReadString('\n')

		nums := strings.Split(input, ",")
		fmt.Println("Input string nums")
		fmt.Println(nums)

		for _, val := range nums {
			fmt.Println("string num ele: ")
			fmt.Println(val)
			fmt.Println("string num Type is ")
			fmt.Println(reflect.TypeOf(val))

			if val == "X" {
				fmt.Println("EXITING..........")
				os.Exit(3)
			}

			//slice_i, _ := strconv.Atoi(val)
			slice_i, _ := strconv.ParseInt(val, 10, 0)
			fmt.Println("to be added into slice int: ")
			fmt.Println(slice_i)
			sli1 = append(sli1, int(slice_i))
		}
		fmt.Println("current Slice is : ")
		fmt.Println(sli1)
		//sli1 = append(sli1, i)

		// fmt.Println("current length is : ")
		// fmt.Println(len(sli1))
		sort.Ints(sli1)
		fmt.Println("current sorted array is : ")

		fmt.Println(sli1)
		fmt.Println("-------------------")

	}
}
