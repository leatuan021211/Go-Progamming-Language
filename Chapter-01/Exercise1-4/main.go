package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	dup_files := make(map[string]string)
	files := os.Args[1:]
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Duplicate: %v\n", err)
			continue
		}
		countLines(f, counter, dup_files)
		f.Close()
	}
	for line, count := range counter {
		if count > 1 {
			fmt.Printf("%d\t%s\t%s\n", line, count, dup_files[line])
		}
	}

}

func countLines(f *os.File, counter map[string]int, dup_files map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counter[line]++
		if counter[line] > 1 {
			if _, exists := dup_files[line]; !exists {
				// Record the filename where the duplicate was found
				dup_files[line] += f.Name() + ","
			}
		}
	}
}
