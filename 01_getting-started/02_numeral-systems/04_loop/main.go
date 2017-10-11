package main

import "fmt"

func main() {
	for x := 0; x <= 20; x++ {
		fmt.Printf("%d \t %b \t %#X \n", x, x, x)
	}
}
