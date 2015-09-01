/**
Methods continued
You can declare a method on any type that is declared in your package, not just struct types.

However, you cannot define a method on a type from another package (including built in types).
 */
package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type Age int

func (age Age) By(v int) int {
	return int(age) * v + v
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	a := Age(12)
	fmt.Printf("\nAge=%d",a.By(2))
}
