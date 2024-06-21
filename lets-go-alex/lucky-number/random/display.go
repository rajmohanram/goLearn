package random

import (
	"fmt"
)

func DisplayRandomNumber() string {
	n := Number()
	return fmt.Sprintf("The random number is %d!\n", n)
	
}

