package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	b, _ := os.Open(os.Args[1])

	_, _ = io.Copy(os.Stdout, b)
}
