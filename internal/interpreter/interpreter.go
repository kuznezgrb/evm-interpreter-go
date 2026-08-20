package interpreter

import (
	"fmt"
	"math/big"

	"github.com/kuznezgrb/evm-interpreter-go/internal/stack"
)

type Opcode byte
type Operation func() error

const (
	STOP  Opcode = 0x00
	PUSH1 Opcode = 0x60

	POP Opcode = 0x50

	DUP1 Opcode = 0x80
	DUP2 Opcode = 0x81

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
		PUSH1: inter.push1,
		POP:   inter.pop,
		DUP1:  inter.dup1,
		DUP2:  inter.dup2,
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

func (i *Interpreter) push1() error {
	if i.pc+1 >= i.lenCode {
		return fmt.Errorf("no value to add to stack")
	}

	val := new(big.Int).SetBytes(i.code[i.pc+1 : i.pc+2])
	i.stack.Push(val)
	i.pc += 2

	return nil
}

func (i *Interpreter) pop() error {
	defer func() {
		i.pc++
	}()
	i.stack.Pop()

	return nil
}

func (i *Interpreter) dup1() error {
	defer func() {
		i.pc++
	}()

	val := i.stack.PeekN(0)
	i.stack.Push(val)
	return nil
}

func (i *Interpreter) dup2() error {
	defer func() {
		i.pc++
	}()

	val := i.stack.PeekN(1)
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
