package interpreter

import (
	"fmt"
	"math/big"

	"github.com/kuznezgrb/evm-interpreter-go/internal/stack"
)

type Opcode byte
type Operation func() error

const (
	STOP   Opcode = 0x00
	PUSH0  Opcode = 0x5F
	PUSH1  Opcode = 0x60
	PUSH2  Opcode = 0x61
	PUSH3  Opcode = 0x62
	PUSH4  Opcode = 0x63
	PUSH5  Opcode = 0x64
	PUSH6  Opcode = 0x65
	PUSH7  Opcode = 0x66
	PUSH8  Opcode = 0x67
	PUSH9  Opcode = 0x68
	PUSH10 Opcode = 0x69
	PUSH11 Opcode = 0x6A
	PUSH12 Opcode = 0x6B
	PUSH13 Opcode = 0x6C
	PUSH14 Opcode = 0x6D
	PUSH15 Opcode = 0x6E
	PUSH16 Opcode = 0x6F
	PUSH17 Opcode = 0x70
	PUSH18 Opcode = 0x71
	PUSH19 Opcode = 0x72
	PUSH20 Opcode = 0x73
	PUSH21 Opcode = 0x74
	PUSH22 Opcode = 0x75
	PUSH23 Opcode = 0x76
	PUSH24 Opcode = 0x77
	PUSH25 Opcode = 0x78
	PUSH26 Opcode = 0x79
	PUSH27 Opcode = 0x7A
	PUSH28 Opcode = 0x7B
	PUSH29 Opcode = 0x7C
	PUSH30 Opcode = 0x7D
	PUSH31 Opcode = 0x7E
	PUSH32 Opcode = 0x7F

	POP Opcode = 0x50

	DUP1  Opcode = 0x80
	DUP2  Opcode = 0x81
	DUP3  Opcode = 0x82
	DUP4  Opcode = 0x83
	DUP5  Opcode = 0x84
	DUP6  Opcode = 0x85
	DUP7  Opcode = 0x86
	DUP8  Opcode = 0x87
	DUP9  Opcode = 0x88
	DUP10 Opcode = 0x89
	DUP11 Opcode = 0x8A
	DUP12 Opcode = 0x8B
	DUP13 Opcode = 0x8C
	DUP14 Opcode = 0x8D
	DUP15 Opcode = 0x8E
	DUP16 Opcode = 0x8F

	SWAP1  Opcode = 0x90
	SWAP2  Opcode = 0x91
	SWAP3  Opcode = 0x92
	SWAP4  Opcode = 0x93
	SWAP5  Opcode = 0x94
	SWAP6  Opcode = 0x95
	SWAP7  Opcode = 0x96
	SWAP8  Opcode = 0x97
	SWAP9  Opcode = 0x98
	SWAP10 Opcode = 0x99
	SWAP11 Opcode = 0x9A
	SWAP12 Opcode = 0x9B
	SWAP13 Opcode = 0x9C
	SWAP14 Opcode = 0x9D
	SWAP15 Opcode = 0x9E
	SWAP16 Opcode = 0x9F
)

type Interpreter struct {
	stack      *stack.Stack
	pc         int
	code       []byte
	lenCode    int
	operations map[Opcode]Operation
}

func NewInterpreter(stack *stack.Stack, code []byte) *Interpreter {

	inter := Interpreter{
		stack:   stack,
		pc:      0,
		code:    code,
		lenCode: len(code),
	}

	operations := map[Opcode]Operation{
		STOP:  inter.stop,
		PUSH0: inter.push0,
		POP:   inter.pop,
	}

	for op := PUSH1; op <= PUSH32; op++ {
		n := int(op - PUSH1 + 1)
		operations[op] = func() error {
			return inter.push(n)
		}
	}

	for op := DUP1; op <= DUP16; op++ {
		n := int(op - DUP1)
		operations[op] = func() error {
			return inter.dup(n)
		}
	}

	for op := SWAP1; op <= SWAP16; op++ {
		n := int(op - SWAP1 + 1)
		operations[op] = func() error {
			return inter.swap(n)
		}
	}

	inter.operations = operations

	return &inter
}

func (i *Interpreter) Run() {
	for i.pc < i.lenCode {
		if err := i.next(); err != nil {
			panic(err)
		}
	}
}

func (i *Interpreter) next() error {
	opcode := Opcode(i.code[i.pc])
	operation, ok := i.operations[opcode]
	if !ok {
		panic(fmt.Sprintf("unknown opcode: 0x%x", opcode))
	}
	return operation()
}

func (i *Interpreter) stop() error {
	i.pc = i.lenCode
	return nil
}

func (i *Interpreter) push0() error {
	val := new(big.Int).SetBytes([]byte{0x00})
	i.stack.Push(val)
	i.pc++

	return nil
}

func (i *Interpreter) push(s int) error {
	if i.pc+s+1 > i.lenCode {
		return fmt.Errorf("no value to add to stack")
	}

	val := new(big.Int).SetBytes(i.code[i.pc+1 : i.pc+s+1])
	i.stack.Push(val)
	i.pc += s + 1

	return nil
}

func (i *Interpreter) pop() error {
	defer func() {
		i.pc++
	}()
	i.stack.Pop()

	return nil
}

func (i *Interpreter) dup(s int) error {
	defer func() {
		i.pc++
	}()

	val := i.stack.PeekN(s)
	i.stack.Push(val)
	return nil
}

func (i *Interpreter) swap(s int) error {
	defer func() {
		i.pc++
	}()
	i.stack.Swap(s)
	return nil
}
