package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	type Person struct {
		name string
		addr string
		phone string

	}
	p1 := Person{
		name: "joe", 
		addr: "123 st",
		phone: "233647774"
}
	//json
	js, err := json.Marshal(p1)

	var Person p2

	err := json.Unmarshal(js, &p2)

	//ioutil
	dat, e := ioutil.ReadFile("rest.txt")

	//os.Open()
	//os.Close()
	//os.Read()
	//os.Write()
	//os.Create()
}