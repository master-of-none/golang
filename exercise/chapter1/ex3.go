package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, name := range counts {
		if len(name) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(name), line, name)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
}
