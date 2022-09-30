#!/usr/bin/env sh
set -e

echo "=== Add Records ==="
./igbo-go -add "overhead press: 70lbs"
./igbo-go -add "20 minute walk"

echo "=== Retrieve Records ==="
./igbo-go -get 1 | grep "overhead press"
./igbo-go -get 2 | grep "20 minute walk"

echo "=== List Records ==="
./igbo-go -list
./igbo-go -list  | grep "overhead press"
./igbo-go -list  | grep "20 minute walk"