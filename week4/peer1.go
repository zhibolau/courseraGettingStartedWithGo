package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	user := make(map[string]string)
	fmt.Println("Enter your name.")
	if scanner.Scan() {
		name := scanner.Text()
		user["name"] = name
	}
	fmt.Println("Enter your address.")
	if scanner.Scan() {
		address := scanner.Text()
		user["address"] = address
	}
	obj, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatal("Error is: ", err)
		return
	}
	fmt.Println("Json object is ", string(obj))
}
