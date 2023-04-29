package main

import (
	"fmt"
	"os"
)

func main() {
	sep := ""

	for _, arg := range os.Args[1:] {
		sep += arg
		sep += " "
	}
	fmt.Println(sep)
}
