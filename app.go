package main

import "fmt"

func main() {
	fmt.Print(sum(5, 5), "\n")
}

func sum(arg1 float32, arg2 float32) float32 {
	return arg1 + arg2
}
