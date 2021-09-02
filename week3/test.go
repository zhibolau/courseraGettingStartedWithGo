package main

import (
	"fmt"
	"os"
)

func main() {

	//array

	var x [5]int
	x[0] = 2
	fmt.Println(x[1])

	var y [5]int = [5]int{1, 2, 3, 4, 5} // array literal: predefined array
	fmt.Println(y[2])

	for i, v := range y {
		fmt.Println("ind %d, val %d", i, v)
		//fmt.Println(i, v)
	}

	//slice is a window of an array
	//can change size <= array size
	//capacity = start to the end of array
	arr := [...]string{"a", "b", "c", "d"}
	s1 := arr[1:3] //3rd element is not included
	fmt.Println(len(s1), cap(s1))
	//updating slice will change the underlying array
	//slice literal == array literal

	//s2 := []int{1, 2, 3} //empty[] means thisis a slice

	//variable slice
	//vs1 := make([]int, 4)
	//vs2 := make([]int, 19, 30)
	//append can increase array size
	//vs3 := append(vs1, 100)

	//map
	//var idMap map[string]int
	//idMap = make(map[string]int)

	idMap1 := map[string]int{"joe": 123}
	fmt.Println(idMap1["joe"])

	for k, v := range idMap1 {
		fmt.Println(k, v)
	}

	//struct
	//person struct :   name, address, phone
	type Person struct {
		name  string
		addr  string
		phone string
	}

	// var p1 Person
	// p1 = Person{name: "joe", addr: "asdfasdf", phone: "12312312312312"}
	// p2 := new(Person)

	val := "X"
	if val == "X" {
		fmt.Println("EXITING..........")
		os.Exit(3)
	}
}
