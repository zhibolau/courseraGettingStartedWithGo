package main

import "fmt"

func main() {
	var x int = 2
	y := &x
	z := *y
	fmt.Println(x) //this is a single comment
	fmt.Println(y)
	fmt.Println(z)
	/*

	   BLOCK COMMENT

	*/
	var a int
	var b *int
	b = &x
	a = *b
	fmt.Println(a)
	fmt.Println(b)

	c := "3d"
	fmt.Println(IsDigit(c))
}
