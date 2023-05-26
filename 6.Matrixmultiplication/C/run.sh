#!/bin/bash

echo "Exercise 6 - C"
cd "C"
gcc -g -o main main2.c
objdump -d main > disassembly.txt
wc -l disassembly.txt
./main
cd ..