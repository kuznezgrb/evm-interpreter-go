package stack_test

import (
	"math/big"
	"testing"

	"github.com/kuznezgrb/evm-interpreter-go/internal/stack"
)

func TestStack_Push(t *testing.T) {
	stack := stack.NewStack()

	one := big.NewInt(1)
	two := big.NewInt(2)

	stack.Push(one)
	stack.Push(two)
}

func TestStack_Push1024(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("the stack is not overflowed")
		}
	}()

	stack := stack.NewStack()

	val := big.NewInt(1)

	for range 1025 {
		stack.Push(val)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := stack.NewStack()

	one := big.NewInt(1)
	two := big.NewInt(2)

	stack.Push(one)
	stack.Push(two)

	topStack := stack.Pop()

	if two.Cmp(topStack) != 0 {
		t.Errorf("got %v, want %v", two, topStack)
	}

	topStack = stack.Pop()
	if one.Cmp(topStack) != 0 {
		t.Errorf("got %v, want %v", one, topStack)
	}
}

func TestStack_PopEmpty(t *testing.T) {
	stack := stack.NewStack()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, but function did not panic")
		}
	}()

	stack.Pop()
}
