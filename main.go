package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/kuznezgrb/evm-interpreter-go/internal/interpreter"
	"github.com/kuznezgrb/evm-interpreter-go/internal/stack"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("required argument is missing")
		return
	}

	argCode := strings.TrimPrefix(os.Args[1], "0x")

	code, err := hex.DecodeString(argCode)

	if err != nil {
		panic("code cannot be recognized")
	}

	stack := stack.NewStack()
	interp := interpreter.NewInterpreter(stack, code)

	interp.Run()

}
