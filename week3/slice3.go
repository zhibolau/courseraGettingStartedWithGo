package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 3)

	fmt.Print("Enter Numbers: ")
	for 1 > 0 {
		r := bufio.NewReader(os.Stdin)
		line, err := r.ReadString('\n')
		if strings.Contains(line, "X") {
			fmt.Println("Exiting program as X/x or non-int is entered!  ")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		fields := strings.Fields(line)
		for _, f := range fields {
			var temp int
			temp, _ = strconv.Atoi(f)
			ints = append(ints, temp)
			if err != nil {
				fmt.Print(err)
			}
		}

		sort.Ints(ints)
		fmt.Println("ints:", ints)

	}

}
