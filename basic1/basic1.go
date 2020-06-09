package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int { // could be shortened => add(x, y int) int {}
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2
var k, l, m = true, 1, "hello"

func main() {
	fmt.Println("Welcome to the playground")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10)) // executed deterministic, will always give the same number
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi) // exported name begins with capital letter

	// add
	fmt.Println("2 + 5 =", add(2, 5))

	// swap
	a, b := swap("world", "hello")
	fmt.Println(a, b)

	// split
	fmt.Println(split(17))

	fmt.Println(c, python, java)
	fmt.Println(i, j)
	fmt.Println(k, l, m)

	x := 10 // implicit statement
	fmt.Println(x)

	f := float64(i) // type conversion
	fmt.Println(f)

	var q int
	var w float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", q, w, bo, s) // value
	fmt.Printf("%T %T %T %T\n", q, w, bo, s) // type

	const Pi = 3.14 // constant
	fmt.Println(Pi)

}
