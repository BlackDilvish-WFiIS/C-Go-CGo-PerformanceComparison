#!/bin/bash

echo ""
echo "=================="
echo "Exercise 3 FilesIO"

cd "3.FilesIO"
echo "Creating a file with 50000 lines"
seq 1 50000 > input.txt
./C/run.sh
./Go/run.sh
./CGo/run.sh
rm input.txt
cd ..