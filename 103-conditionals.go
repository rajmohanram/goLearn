// if conditional

package main

import (
	"fmt"
)

func main() {
	x := 10
	
	// if
	if x > 5 {
		fmt.Println("x is big")
	}

	// if with logical AND
	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	// if with logical OR
	if x <20  || x < 30 {
		fmt.Println("x is out of range")
	}

	// if and else
	if x > 100 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is not that big")
	}

	a := 11.0
	b := 20.0

	if frac:= a/b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
	

}