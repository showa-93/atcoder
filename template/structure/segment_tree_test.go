package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRMQ(t *testing.T) {
	a := []int{4, 8, 2, 3, 10, 9, 1, 6, 7, 5}
	rmq := NewRMQ(10, 1<<16, Min)
	for i, ai := range a {
		rmq.Update(i, ai)
	}

	assert.Equal(t, 1, rmq.Query(0, 10))
	assert.Equal(t, 3, rmq.Query(3, 6))
	assert.Equal(t, 2, rmq.Query(1, 4))
	assert.Equal(t, 6, rmq.Query(7, 9))
	assert.Equal(t, 1<<16, rmq.Query(10, 11))
}

func TestBIT(t *testing.T) {
	a := []int{5, 3, 7, 9, 6, 4, 1, 2}
	bit := &BIT{
		n:   8,
		bit: make([]int, 1<<16+1),
	}
	for i, ai := range a {
		bit.Add(i+1, ai)
	}

	assert.Equal(t, 5, bit.Sum(1))
	assert.Equal(t, 30, bit.Sum(5))
	assert.Equal(t, 37, bit.Sum(8))
}
