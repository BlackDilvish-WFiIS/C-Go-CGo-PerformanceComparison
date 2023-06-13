#!/bin/bash

echo ""
echo "=================="
echo "Exercise 2 Insertion Sort"

cd "2.InsertionSort"
if [[ ! -f strings.txt ]]; then
    echo "Generating random strings"
    N=10000; LEN=5; for i in $(seq 1 $N); do tr -dc A-Za-z0-9 </dev/urandom | head -c $LEN ; echo; done > strings.txt
fi
if [[ ! -f ints.txt ]]; then
    echo "Generating random ints"
    N=10000; for i in $(seq 1 $N); do echo $RANDOM; done > ints.txt
fi

./C/run.sh
./Go/run.sh
./CGo/run.sh
cd ..