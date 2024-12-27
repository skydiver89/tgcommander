package main

import (
	"fmt"
	"log"
)

func main() {
	var config Config
	err := config.load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", config)
}
