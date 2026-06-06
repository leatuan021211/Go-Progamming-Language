package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for idx, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(idx) + " " + arg
		sep = " "
	}
	fmt.Println(s)
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	duration = time.Since(start)
	fmt.Println(duration)
}
