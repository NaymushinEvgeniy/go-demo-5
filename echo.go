package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		fmt.Println(args[i])
	}
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		sep = " "
	}
	fmt.Println(s)
}
