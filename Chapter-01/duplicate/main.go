package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	counter := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counter[input.Text()]++
// 	}
// 	for line, n := range counter {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

func main() {
	counter := make(map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Duplicate: %v\n", err)
			continue
		}
		countLines(f, counter)
		f.Close()
	}
	for line, count := range counter {
		if count > 1 {
			fmt.Printf("%d\t%s\n", line, count)
		}
	}

}

func countLines(f *os.File, counter map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counter[input.Text()]++
	}
}
