package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	maxLength = 20
)

type Person struct {
	fname string
	lname string
}

func (n *Person) Set(firstName string, lastName string) {
	var rs []rune

	n.fname = firstName
	if len(firstName) > maxLength {
		rs = []rune(firstName)
		n.fname = string(rs[:maxLength])
	}

	n.lname = lastName
	if len(lastName) > maxLength {
		rs = []rune(lastName)
		n.lname = string(rs[:maxLength])
	}
}

func main() {

	/*
		examples to write in names.txt:
			SMorr5678901234567890AAAAA Morr
			L Roll
			M Max
			M Morr5678901234567890AAAAA
	*/
	fmt.Println("Enter the file name - path:")
	fileName := ""
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil {
		log.Fatal("Error to reading file name: ", err)
		return
	}

	file, _ := os.Open(fileName)
	if err != nil {
		log.Fatal("error opening file. err is ", err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	txtlines := []string{}

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	nameSlice := make([]Person, 0, len(txtlines))
	l := &Person{}

	for _, v := range txtlines {
		someString := v
		words := strings.Fields(someString)
		l.Set(words[0], words[1])
		nameSlice = append(nameSlice, *l)
	}

	for _, j := range nameSlice {
		fmt.Println(j.fname, j.lname)
	}
}
