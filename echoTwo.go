package main

import (
	"fmt"
	"os"
)

func main() {
	sep, s := " ", ""
	fmt.Printf("Name of progs is %s\n", os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Printf("Index of element is %d, value is %s\n", i, arg)
		s += arg + sep
	}
	fmt.Println(s)
}
