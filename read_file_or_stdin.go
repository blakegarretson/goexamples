package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		processLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				continue
			}
			processLines(f)
			f.Close()
		}
	}
}

func processLines(f *os.File) {
	input := bufio.NewScanner(f)
	linenum := 1
	for input.Scan() {
		fmt.Printf("%d %s\n", linenum, input.Text())
		linenum++
	}
	// Should be checking for errors from input.Err()
}
