#!/bin/bash

echo ""
echo "=================="
echo "Exercise 3 FilesIO"

cd "3.FilesIO"
seq 1 10000 > input.txt
./C/run.sh
./Go/run.sh
./CGo/run.sh
cd ..