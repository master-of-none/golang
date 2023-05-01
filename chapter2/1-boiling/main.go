package main

import "fmt"

const boilingF = 212.0

func main() {
	f := boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %gºF or %gºC\n", f, c)
}
