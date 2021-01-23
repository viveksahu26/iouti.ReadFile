package main

import (
	"fmt"
	"io/ioutil"
	//"log"
)

func check(e error) {
    if e != nil {
        panic(e)
    } 
}
func main() {
  //WriteFile writes data to a file named by filename. If the file does not exist,
  //WriteFile creates it with permissions
  data:=[]byte("All the data I wish to write to a file name data.txt")

  err:= ioutil.WriteFile("data.txt", data, 0777) 
  check(err)
  fmt.Println("Successfully written data into data.txt\n")
  
  //ReadFile bring file into memory and read whole
  //file at once and return in form of byte to buf.
	content, err := ioutil.ReadFile("data.txt")
	check(err)
  // Converting bytes of data into strings.
	fmt.Println("Now Reading file Data.txt :-\n"+(string(content)))

}