package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" // only available in functions scope
	fmt.Println(f)

	fmt.Println(s)

	const n = 500000000

	const D = 3e20 / n
	fmt.Println(D)

	fmt.Println(int64(D))

	fmt.Println(math.Sin(n))
}
