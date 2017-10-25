package main

import "fmt"

var x int = 42

// scope of x since it is not in curly braces is of the whole package
// i.e. package level scope

func main() {
	fmt.Println(x)
	poo()
	loo()
	boo()
}

func poo() {
	fmt.Println("1")
	boo()
}

func boo() {
	fmt.Println("2")
	y := 17
	fmt.Println(y)

}

func loo() {
	fmt.Println("3")
}

//between the braces block level scope enclosed in larger scope area
