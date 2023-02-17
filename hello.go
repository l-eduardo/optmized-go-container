package main

import (
	"os"
)

func Println() (n int, err error) {
	return os.Stdout.WriteString("hello world\n")
}

func main() {
	Println()
}
