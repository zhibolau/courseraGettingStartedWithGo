/*

Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file
 has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname
for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct
which contains the first and last names found in the file. Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice containing one struct for
each line in the file. After reading all lines from the file, your program should iterate through your slice
 of structs and print the first and last names found in each struct.



*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

const (
	maxLength = 20
)

func main() {
	type Name struct {
		fname, lname string
	}

	fmt.Println("Please enter file full name and file needs to be at same folder as this go file: ")
	r := bufio.NewReader(os.Stdin)
	fileName, err1 := r.ReadString('\n')
	check(err1)

	//fullFileName := string(fileName) + ".txt"
	//fmt.Println(fileName[:len(fileName)-1])
	//fmt.Println(reflect.TypeOf(fileName))
	path := fileName[:len(fileName)-1]
	// dat, err2 := os.ReadFile(path)
	// check(err2)

	//fmt.Print(string(dat))
	res := make([]Name, 0)
	buf, err2 := os.Open(path)
	snl := bufio.NewScanner(buf)
	for snl.Scan() {

		name_str := strings.Split(snl.Text(), " ")
		n1 := Name{fname: name_str[0]}
		n1.lname = name_str[1]
		//fmt.Println(snl.Text())
		res = append(res, n1)
		//fmt.Println(n1)
	}
	check(err2)
	fmt.Println("printing Name struc: --------->")
	for _, v := range res {
		//fmt.Print(v)
		fmt.Print("fname: " + v.fname)
		fmt.Print(" ")
		fmt.Print("lname: " + v.lname)
		fmt.Println()
	}
}
