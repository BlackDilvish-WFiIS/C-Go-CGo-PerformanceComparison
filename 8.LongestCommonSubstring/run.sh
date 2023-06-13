#!/bin/bash

echo ""
echo "=================="
echo "Exercise 8 Longest Common Substring"

cd "8.LongestCommonSubstring"
if [[ ! -f strings.txt ]]; then
    echo "Generating random string"
    N=2; LEN=10000; for i in $(seq 1 $N); do tr -dc A-Za-z0-9 </dev/urandom | head -c $LEN ; echo; done > strings.txt
fi

./C/run.sh
./Go/run.sh
./CGo/run.sh
cd ..