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
	loo()
}

func boo() {
	fmt.Println("2")
}

func loo(){
	fmt.Println("3")
}
