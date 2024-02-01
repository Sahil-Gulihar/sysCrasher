package main

import (
	"io/ioutil"
	"log"
	"strings"
)


func main() {

	msg := []byte (strings.Repeat("Hello",1024*1024*500))
	err := ioutil.WriteFile("hello", msg , 0644)

	if err != nil{
		log.Fatal(err)
	}

}
