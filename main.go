package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Too few arguments!\nUsage: hiddenprint FILE\n")
		os.Exit(1)
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Encountered error: %v\n", err.Error())
		os.Exit(1)
	}
	// Quick & dirty way to remove the automatically added "" by Sprintf
	fmt.Println(strings.TrimSuffix(strings.TrimPrefix(fmt.Sprintf("%q", string(file)), "\""), "\""))
}
