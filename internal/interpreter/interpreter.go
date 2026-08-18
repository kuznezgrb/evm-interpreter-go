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
