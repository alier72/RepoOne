package main

import (
	"fmt"
	"math/rand"
)

func test3(sum int) int {
   	var x,y int
	x=sum/2
	y=x-4
	var z string
	_ = z
	sum=sum+y
   return sum
}

func mod1() {
	fmt.Println(19)
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	mod1()
	fmt.Println(test3(34))
}
