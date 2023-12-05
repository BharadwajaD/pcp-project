#!/bin/bash

# Array of numbers
numbers=(1 2 4 6 8 16 32)

# Iterate over the array
for a in "${numbers[@]}"; do
    echo $a
    THREAD_COUNT="$a" go run main.go -input ./datasets/dataset2.txt -algo boruvaka
done

