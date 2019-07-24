package golib

import (
	"fmt"
	"math"
	"sync/atomic"
)

func NewRoundRobin(bound int) *RoundRobin {
	if bound <= 0 {
		panic(fmt.Sprintf("invalid bound: %d", bound))
	}
	return &RoundRobin{
		bound: uint64(bound),
		index: math.MaxUint64,
	}
}

type RoundRobin struct {
	bound uint64
	index uint64
}

func (rr *RoundRobin) Next() int {
	index := atomic.AddUint64(&rr.index, 1)
	return int(index % rr.bound)
}
