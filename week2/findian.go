/*
Write a program which prompts the user to enter a string. The program searches through the entered
 string for the characters ‘i’, ‘a’, and ‘n’.

 The program should print “Found!” if the entered string
 starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.

 The program should print “Not Found!” otherwise.

Examples: The program should print
“Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
 The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var targetStr string

	// fmt.Println("plz enter a string:")
	// _, err := fmt.Scan(&targetStr)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println("plz enter a string:")
	targetStr := bufio.NewReader(os.Stdin)
	input, _ := targetStr.ReadString('\n')

	fmt.Println("this is what you entered: " + input)

	//lower_str_space := strings.ToLower(input)
	lower_str := strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(input), "\n", ""), " ", "")
	//lower_str := strings.ReplaceAll(lower_str_space, " ", "")

	//fmt.Println("lower str without spaces is : " + lower_str)

	begin_i := strings.HasPrefix(lower_str, "i")
	//fmt.Println(begin_i)
	end_n := strings.HasSuffix(lower_str, "n")
	//fmt.Println(end_n)
	has_a := strings.Contains(lower_str, "a")
	// fmt.Println(has_a)

	if begin_i && end_n && has_a {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
