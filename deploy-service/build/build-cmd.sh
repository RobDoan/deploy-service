#!/bin/bash

set -euo pipefail

# Project root directory
ROOT_DIR=$(git rev-parse --show-toplevel)

PROJECT_ROOT="$ROOT_DIR/deploy-service"

# Destination folder for built binaries
OUTPUT_DIR="$PROJECT_ROOT/bin"

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Find all cmd packages within the project
CMD_PACKAGES=$(find "$PROJECT_ROOT/cmd" -type d -mindepth 1 -maxdepth 1 -exec basename {} \;)

# Build each cmd package and move the binary to the output directory
for CMD_PACKAGE in $CMD_PACKAGES; do
    go build -o "$OUTPUT_DIR/$CMD_PACKAGE" "$PROJECT_ROOT/cmd/$CMD_PACKAGE"
done

# Print a confirmation message
echo "All cmd apps built successfully and moved to $OUTPUT_DIR"