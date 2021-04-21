package main

import (
	fibo "fibonacci/pkg/fibo"
	validator "fibonacci/pkg/validator"
	"fmt"
)

func main() {
	var minValue, maxValue = 0, 20
	var num int = 20
	var isValid bool = validator.InRange(num, minValue, maxValue)
	if !isValid {
		fmt.Printf("%d - wrong index.Index number mast be in range from %d to %d.\n", num, minValue, maxValue)
		return
	}

	fmt.Printf("Value of the fibonacci number at the index %d is: %d \n", num, fibo.Calc(num))
}
