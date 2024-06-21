package main

import (
	"fmt"
	"os"
)

func main() {
	sep, s := " ", ""
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}
