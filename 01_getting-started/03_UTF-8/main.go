package main

import "fmt"

func main() {
	for x := 65; x <= 122; x++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", x, x, x, x)
		// %d returns, decimals, %b returns binary, %X returns hex %q return UTF
	}
}
