package interpreter

import "github.com/kuznezgrb/evm-interpreter-go/internal/stack"

type Interpreter struct {
	stack *stack.Stack
}

func NewInterpreter(stack *stack.Stack) *Interpreter {
	return &Interpreter{
		stack: stack,
	}
}

func (i *Interpreter) Run(code []byte) {

}
