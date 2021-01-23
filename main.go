package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
  //ReadFile bring file into memory and read whole
  //file at once and return in form of byte to buf.
	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
  // Converting bytes of data into strings.
	fmt.Println((string(content)))

}