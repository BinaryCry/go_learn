package main

import "fmt"
import (
	"math/rand"
	"math"
)

var(
	str string = "HelloWorld"
	n uint8 = 1
	x int8 = 15
)

func main() {
	fmt.Println(len(str))
	fmt.Println(str[n])
	fmt.Println(n % 5)

	var inp float64
	fmt.Print("Enter F temp: ")
	fmt.Scanf("%f", &inp)
	var res float64
	res = (inp - 32) * 5 / 9
	fmt.Println("Celsium:", math.Ceil(res))
	fmt.Println(rand.Int())
}


