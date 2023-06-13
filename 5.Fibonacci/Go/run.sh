#!/bin/bash

echo "Excercise 5 - Go"
cd "Go"
go build -gcflags "-N -l" -o main main.go
go tool objdump -S main > disassembly.txt
wc -l disassembly.txt
time ./main
cd ..
