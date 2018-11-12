package main

import (
	"github.com/alphamu/go-collections/collections"
)

func main() {
	var factorOf3 = collections.Condition{3, "Fizz"}
	var factorOf5 = collections.Condition{5, "Buzz"}
	var fizzBuzz = &collections.FizzBuzz{}
	fizzBuzz.Tests = make([]collections.Condition, 3)
	fizzBuzz.Tests[0] = factorOf3
	fizzBuzz.Tests[1] = factorOf5

	for i := 0.0; i < 100.0; i++ {
		if !fizzBuzz.Eval(i) {
			fizzBuzz.Fallback(i)
		}
	}
}