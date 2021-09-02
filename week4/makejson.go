/*
Write a program which prompts the user to first enter a name, and
then enter an address. Your program should create a map and add the name
and address to the map using the keys “name” and “address”, respectively.
 Your program should use Marshal() to create a JSON object from the map,
 and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	m := make(map[string]string)

	fmt.Println("Enter name and address: ")
	res2 := make([]string, 2)
	for i := 0; i < 2; i++ {
		r := bufio.NewReader(os.Stdin)
		line, err := r.ReadString('\n')
		res2 = append(res2, line)
		if err != nil {
			log.Fatal(err)
		}
	}
	res := res2[2:]

	m["name"] = res[0]
	m["address"] = res[1]

	// fmt.Println("person ")
	// fmt.Printf("%+v\n", m)
	js, err1 := json.Marshal(m)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("JSON COMING: ")
	fmt.Println(string(js))
}
