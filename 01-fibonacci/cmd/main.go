package main

import (
	"fibonacci/pkg/fibo"
	"fibonacci/pkg/validator"
	"fmt"
)

func main() {
	var min, max = 0, 20
	var num int = 6
	var isValid bool = validator.InRange(num, min, max)

	if !isValid {
		fmt.Printf("%d - wrong index.Index number mast be in range from %d to %d.\n", num, min, max)

		return
	}

	fmt.Printf("Value of the fibonacci number at the index %d is: %d \n", num, fibo.Num(num))
}
