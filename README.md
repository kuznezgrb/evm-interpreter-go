# evm-interpreter-go

A minimal Ethereum Virtual Machine (EVM) bytecode interpreter written in Go, built from scratch.

## Project layout

```
.
├── main.go                          # CLI entry point
├── internal/
│   ├── interpreter/
│   │   └── interpreter.go           # Interpreter that runs EVM bytecode
│   └── stack/
│       ├── stack.go                 # 1024-slot EVM stack (big.Int based)
│       └── stack_test.go
└── go.mod
```

## Requirements

- Go 1.26+

## Usage

```bash
go run . <hex-encoded-bytecode>
```

Example:

```bash
go run . 0x6001600201
```

The `0x` prefix is optional.

## Running tests

```bash
go test ./...
```

## License

MIT
