#!/bin/bash

# Array of numbers
numbers=(1 2 4 8 16)

# Iterate over the array
for a in "${numbers[@]}"; do
    echo $a
    THREAD_COUNT="$a" go run main.go -input ./datasets/input.txt
done

