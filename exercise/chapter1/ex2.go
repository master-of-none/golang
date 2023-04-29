package main

import (
	"fmt"
	"os"
)

func main() {
	// sep := ""
	i := 1
	for _, args := range os.Args[1:] {
		// sep = args
		fmt.Println("Index:", i, "Args", args)
		i++
	}
}
