#!/bin/bash

echo "Exercise 1 - Go"
cd "Go"
go build -gcflags "-N -l" -o main
go tool objdump -S main > disassembly.txt
wc -l disassembly.txt
time ./main > /dev/null
cd ..