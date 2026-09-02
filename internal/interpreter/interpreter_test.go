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

func TestInterpreter_Push0(t *testing.T) {
	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, []byte{0x5f})
	interp.Run()

	valReference := new(big.Int).SetBytes([]byte{0x00})
	stackVal := stack.Pop()

	if valReference.Cmp(stackVal) != 0 {
		t.Errorf("got %v, want %v", stackVal, valReference)
	}
}

func pushBytecode(n int, opcode interpreter.Opcode) (code []byte, payload []byte) {
	payload = make([]byte, n)
	for v := 1; v <= n; v++ {
		payload[v-1] = byte(v)
	}

	code = append([]byte{byte(opcode)}, payload...)
	return code, payload
}

func TestInterpreter_Push(t *testing.T) {
	tests := []struct {
		name   int
		n      int
		opcode interpreter.Opcode
	}{
		{1, 1, interpreter.PUSH1},
		{2, 2, interpreter.PUSH2},
		{3, 3, interpreter.PUSH3},
		{4, 4, interpreter.PUSH4},
		{5, 5, interpreter.PUSH5},
		{6, 6, interpreter.PUSH6},
		{7, 7, interpreter.PUSH7},
		{8, 8, interpreter.PUSH8},
		{9, 9, interpreter.PUSH9},
		{10, 10, interpreter.PUSH10},
		{11, 11, interpreter.PUSH11},
		{12, 12, interpreter.PUSH12},
		{13, 13, interpreter.PUSH13},
		{14, 14, interpreter.PUSH14},
		{15, 15, interpreter.PUSH15},
		{16, 16, interpreter.PUSH16},
		{17, 17, interpreter.PUSH17},
		{18, 18, interpreter.PUSH18},
		{19, 19, interpreter.PUSH19},
		{20, 20, interpreter.PUSH20},
		{21, 21, interpreter.PUSH21},
		{22, 22, interpreter.PUSH22},
		{23, 23, interpreter.PUSH23},
		{24, 24, interpreter.PUSH24},
		{25, 25, interpreter.PUSH25},
		{26, 26, interpreter.PUSH26},
		{27, 27, interpreter.PUSH27},
		{28, 28, interpreter.PUSH28},
		{29, 29, interpreter.PUSH29},
		{30, 30, interpreter.PUSH30},
		{31, 31, interpreter.PUSH31},
		{32, 32, interpreter.PUSH32},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("PUSH%d", tt.name), func(t *testing.T) {
			stack := stack.NewStack()
			code, payload := pushBytecode(tt.n, tt.opcode)
			interp := interpreter.NewInterpreter(stack, code)
			interp.Run()

			valReference := new(big.Int).SetBytes(payload)
			stackVal := stack.Pop()

			if valReference.Cmp(stackVal) != 0 {
				t.Errorf("got %v, want %v", stackVal, valReference)
			}
		})
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

func dupBytecode(n int, opcode interpreter.Opcode) []byte {
	code := make([]byte, 0, n*2+1)
	for v := 1; v <= n; v++ {
		code = append(code, byte(interpreter.PUSH1), byte(v))
	}
	code = append(code, byte(opcode))
	return code
}

func TestInterpreter_Dup(t *testing.T) {
	tests := []struct {
		name   int
		n      int
		opcode interpreter.Opcode
	}{
		{1, 1, interpreter.DUP1},
		{2, 2, interpreter.DUP2},
		{3, 3, interpreter.DUP3},
		{4, 4, interpreter.DUP4},
		{5, 5, interpreter.DUP5},
		{6, 6, interpreter.DUP6},
		{7, 7, interpreter.DUP7},
		{8, 8, interpreter.DUP8},
		{9, 9, interpreter.DUP9},
		{10, 10, interpreter.DUP10},
		{11, 11, interpreter.DUP11},
		{12, 12, interpreter.DUP12},
		{13, 13, interpreter.DUP13},
		{14, 14, interpreter.DUP14},
		{15, 15, interpreter.DUP15},
		{16, 16, interpreter.DUP16},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("DUP%d", tt.name), func(t *testing.T) {
			stack := stack.NewStack()
			interp := interpreter.NewInterpreter(stack, dupBytecode(tt.n, tt.opcode))
			interp.Run()

			valReference := new(big.Int).SetBytes([]byte{0x01})
			stackVal := stack.Pop()

			if valReference.Cmp(stackVal) != 0 {
				t.Errorf("got %v, want %v", stackVal, valReference)
			}
		})
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
