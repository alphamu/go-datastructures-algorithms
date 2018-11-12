package collections

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestQueuePut(t *testing.T) {
	queue := &Queue{}
	// Insert in order 10 -> 0
	// This should result in a queue that pops in the same order
	for i := 10; i >= 0; i-- {
		queue.push(i)
	}
	queue.printQueue()
	fmt.Println()
	// Should pop in order 10 -> 0
	for i := 10; i >= 0; i-- {
		pop := queue.pop()
		fmt.Print(pop, ",")
		assert.True(t, i == pop, fmt.Sprintf("%v should be %v", i, pop))
	}


}