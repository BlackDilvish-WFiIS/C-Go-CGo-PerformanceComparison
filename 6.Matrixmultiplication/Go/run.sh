#!/bin/bash

echo "Exercise 6 - Go"
cd "Go"
go build -gcflags "-N -l" -o main
go tool objdump -S main > disassembly.txt
wc -l disassembly.txt
./main
cd ..