package collections

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)


func TestPutAndGetTree(t *testing.T) {
	bt := BinaryTree{}
	
	for i := 0; i < 20; i++ {
		rand := random(100, 1000)
		fmt.Print(rand, ",")
		bt.Put(rand)
	}
	asc := bt.GetAscending()
	for i:= 1; i < len(asc); i++ {
		assert.True(t, asc[i].(int) > asc[i-1].(int), fmt.Sprintf("%v: %v should be greater than %v", i, asc[i], asc[i-1]))
	}

	dsc := bt.GetDescending()
	for i:= 1; i < len(dsc); i++ {
		assert.True(t, dsc[i].(int) < dsc[i-1].(int), fmt.Sprintf("%v: %v should be less than %v", i, dsc[i], dsc[i-1]))
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}
