package interpreter_test

import (
	"fmt"
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

func swapBytecode(n int, opcode interpreter.Opcode) []byte {
	code := make([]byte, 0, (n+1)*2+1)
	for v := 1; v <= n+1; v++ {
		code = append(code, byte(interpreter.PUSH1), byte(v))
	}
	code = append(code, byte(opcode))
	return code
}

func TestInterpreter_Swap(t *testing.T) {
	tests := []struct {
		name   int
		n      int
		opcode interpreter.Opcode
	}{
		{1, 1, interpreter.SWAP1},
		{2, 2, interpreter.SWAP2},
		{3, 3, interpreter.SWAP3},
		{4, 4, interpreter.SWAP4},
		{5, 5, interpreter.SWAP5},
		{6, 6, interpreter.SWAP6},
		{7, 7, interpreter.SWAP7},
		{8, 8, interpreter.SWAP8},
		{9, 9, interpreter.SWAP9},
		{10, 10, interpreter.SWAP10},
		{11, 11, interpreter.SWAP11},
		{12, 12, interpreter.SWAP12},
		{13, 13, interpreter.SWAP13},
		{14, 14, interpreter.SWAP14},
		{15, 15, interpreter.SWAP15},
		{16, 16, interpreter.SWAP16},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("SWAP%d", tt.name), func(t *testing.T) {
			stack := stack.NewStack()
			interp := interpreter.NewInterpreter(stack, swapBytecode(tt.n, tt.opcode))
			interp.Run()

			valReference := new(big.Int).SetBytes([]byte{0x01})
			stackVal := stack.Pop()

			if valReference.Cmp(stackVal) != 0 {
				t.Errorf("got %v, want %v", stackVal, valReference)
			}
		})
	}
}
