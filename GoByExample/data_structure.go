package main

import (
	"fmt"
	"slices"
)

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println("len:", len(b))

	b = [...]int{1, 2, 3, 4, 5}

	// [...]int is Go syntax that tells the compiler:
	// “Create an array of type int, and let the compiler figure out the length from the number of elements provided.”
	fmt.Println("dcl:", b)
	fmt.Println("len:", len(b))

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	// twoD is a 2-dimensional array with 2 rows and 3 columns.
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

	// Slices are more flexible than arrays.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println("sl0:", s[0])

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twod := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twod[i] = make([]int, innerLen)
		for j := range innerLen {
			twod[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twod)
}
