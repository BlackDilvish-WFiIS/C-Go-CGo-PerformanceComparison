#!/bin/bash

echo "Exercise 8 - C (Optimization 0)"
cd "C"
gcc -g -O0 -o main_O0 main.c
objdump -d main_O0 > disassembly_O0.txt
wc -l disassembly_O0.txt
time ./main_O0

echo "Exercise 8 - C (Optimization 3)"
gcc -g -O3 -o main_O3 main.c
objdump -d main_O3 > disassembly_O3.txt
wc -l disassembly_O3.txt
time ./main_O3
cd ..

