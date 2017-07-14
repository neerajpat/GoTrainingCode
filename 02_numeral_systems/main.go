package main

import "fmt"

func main() {

	for i := 0; i < 123; i++ {
		// print in decimal - binary - hexadecimal  - octal - UTF8
		fmt.Printf("%d \t %b \t %#x \t %#o \t %q \n", i, i, i, i, i)
	}
}
