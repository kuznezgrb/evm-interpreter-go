package stack

import "math/big"

const stackSize = 1024

type Stack struct {
	storage [stackSize]*big.Int
	top     int
}

func NewStack() *Stack {
	return &Stack{
		top: -1,
	}
}

func (s *Stack) Push(val *big.Int) {
	s.top++
	if s.top >= stackSize {
		panic("stack overflow")
	}
	copyVal := new(big.Int).Set(val)
	s.storage[s.top] = copyVal
}

func (s *Stack) Pop() *big.Int {
	if s.top == -1 {
		panic("Attempt to get an element of an empty stack")
	}

	defer func() {
		s.top--
	}()

	return s.storage[s.top]
}

func (s *Stack) PeekN(n int) *big.Int {
	if 0 > n || n > s.top {
		panic("there is no such element in the stack")
	}

	return s.storage[s.top-n]
}
