package main

import (
	"fmt"
	"lucky-number.blueonset.com/random"
	// Import the color package.
    "github.com/fatih/color"
)

func main() {
	
	// Print random number
	num := random.DisplayRandomNumber()
	fmt.Println(num)
	
	// Use it to print the message in green.
    green := color.New(color.FgGreen)
	green.Printf(num)

}
