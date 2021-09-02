package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "names.txt"

	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		fmt.Println(snl.Text())
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
}
