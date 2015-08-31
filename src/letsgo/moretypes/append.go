/**
Adding elements to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the Slices: usage and internals article.)
 */
package main

import "fmt"

func main() {
	var a []int
	printSlice2("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice2("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice2("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice2("a", a)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
