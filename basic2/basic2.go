package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 { // if statement
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // like for, if statement can also start with a short init // variales declared in init also available inside else block
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // v availabe here
	}

	return lim
}

func main() {
	sum1 := 0
	for i := 0; i < 10; i++ { // init, condition, post // init and post are optional
		sum1 += i
	}
	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 1000 { // for with condition only
		sum2 += sum2
	}
	fmt.Println(sum2)

	// sqrt
	fmt.Println(sqrt(2), sqrt(-4))

	// pow
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

	// switch statement // evaluate from top to bottom, stopping when a case succeeds
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch without condition is the same as switch true // useful to write long if-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// defer statement // defers the execution of a function until the surrounding function returns
	// defer pushed onto a stack // LIFO
	defer fmt.Println("!")
	defer fmt.Println("world")

	fmt.Println("Hello")
}
