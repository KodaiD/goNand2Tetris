# Assembler written in Go for Nand2Tetris

I write a simple assembler in Go for nand2tetris. Almost all the function names follow the instructions in the book.

```
$ go build -v -o Assembler
$ ./Assembler File.asm
```

## Instruction set (Function)

- Store data in memory
- Arithmetic operation
- Logical operation
- Fetch and execute the instruction at the specified location

## Feature

- The output binary code runs on the hardware simulator prepared by the textbook.
- Two parses due to the presence of user-defined symbols
