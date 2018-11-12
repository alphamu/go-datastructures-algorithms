package collections

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestStackPut(t *testing.T) {
	stack := &Stack{}
	// Insert in order 10 -> 0
	// this will result in a stack 0 -> 10
	for i := 10; i >= 0; i-- {
		stack.push(i)
	}
	stack.printStack()
	fmt.Println()
	// Should pop in order 0 -> 10
	for i := 0; i <= 10; i++ {
		pop := stack.pop()
		fmt.Print(pop, ",")
		assert.True(t, i == pop, fmt.Sprintf("%v should be %v", i, pop))
	}


}