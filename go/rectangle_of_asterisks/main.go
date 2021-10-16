package main

import (
	"fmt"
	"strings"
)

func main() {
	// print a 5 (row) by 10 (col) grid of asterisks
	str := strings.Repeat("*", 10)
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
}
