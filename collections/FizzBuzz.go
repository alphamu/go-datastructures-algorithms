package collections

import (
	"fmt"
	"math"
)

type IFizzBuzz interface {
	Eval(num float64) bool
	Fallback(num float64) bool
}

type Condition struct {
	FactorOf  float64
	Print     string
}

type FizzBuzz struct {
	Tests []Condition

	IFizzBuzz
}

// returns true if handled, else false
func (t *FizzBuzz) Eval(num float64) bool {
	if num == 0 {
		return false
	}

	var result = ""
	for _, t := range t.Tests {
		if math.Mod(num, t.FactorOf) == 0 {
			result = result + t.Print
		}
	}
	if len(result) > 0 {
		fmt.Println(result)
		return true
	}

	return false
}

func (t *FizzBuzz) Fallback(num float64) bool {
	fmt.Println(num)
	return true
}