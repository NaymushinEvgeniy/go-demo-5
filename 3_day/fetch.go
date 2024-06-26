package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		textBody, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("Content is %s", string(textBody))
	}
}
