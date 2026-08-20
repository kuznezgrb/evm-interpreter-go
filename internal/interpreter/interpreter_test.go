package interpreter_test

import (
	"math/big"
	"testing"

	"github.com/kuznezgrb/evm-interpreter-go/internal/interpreter"
	"github.com/kuznezgrb/evm-interpreter-go/internal/stack"
)

func TestInterpreter_Stop(t *testing.T) {
	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, []byte{0x00})
	interp.Run()
}

func TestInterpreter_Push1(t *testing.T) {
	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, []byte{0x60, 0x2a, 0x00})
	interp.Run()

	valReference := new(big.Int).SetBytes([]byte{0x2a})

	stackVal := stack.Pop()

	if valReference.Cmp(stackVal) != 0 {
		t.Errorf("got %v, want %v", stackVal, valReference)
	}
}

func TestInterpreter_Pop(t *testing.T) {
	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, []byte{0x60, 0x01, 0x60, 0x02, 0x50})
	interp.Run()

	valReference := new(big.Int).SetBytes([]byte{0x01})
	stackVal := stack.Pop()

	if valReference.Cmp(stackVal) != 0 {
		t.Errorf("got %v, want %v", stackVal, valReference)
	}
}

func TestInterpreter_Dup1(t *testing.T) {
	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, []byte{0x60, 0x01, 0x60, 0x02, 0x80})
	interp.Run()

	valReference := new(big.Int).SetBytes([]byte{0x02})
	stackVal := stack.Pop()

	if valReference.Cmp(stackVal) != 0 {
		t.Errorf("got %v, want %v", stackVal, valReference)
	}
}

func TestInterpreter_Dup2(t *testing.T) {
	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, []byte{0x60, 0x01, 0x60, 0x02, 0x81})
	interp.Run()

	valReference := new(big.Int).SetBytes([]byte{0x01})
	stackVal := stack.Pop()

	if valReference.Cmp(stackVal) != 0 {
		t.Errorf("got %v, want %v", stackVal, valReference)
	}
}
