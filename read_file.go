package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			//process file line by line
			fmt.Printf("%s\n", line)
		}
	}
}
