#!/usr/bin/env sh
set -e

echo "=== Add Records ==="
./igbo-go -create ""

echo "=== Retrieve Records ==="
./igbo-go -retrieve ""

echo "=== Update Record ==="
./igbo-go -update ""

echo "=== Retrieve Records ==="
./igbo-go -retrieve ""

echo "=== Delete Record ==="
./igbo-go -delete ""

echo "=== Retrieve Records ==="
./igbo-go -retrieve ""

