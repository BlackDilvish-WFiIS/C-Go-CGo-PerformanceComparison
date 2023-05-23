#!/bin/bash

echo "Exercise 3 - C"
cd "C"
gcc -g -o main main.c
objdump -d main > disassembly.txt
wc -l disassembly.txt
time ./main > /dev/null
cd ..