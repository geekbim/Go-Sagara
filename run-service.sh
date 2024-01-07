#!/bin/sh

ENV_FILE=".env"
MAIN_FILE="main.go"

# Check if the .env file exists
if [ ! -f "$ENV_FILE" ]; then
  echo "Error: $ENV_FILE not found."
  exit 1
fi

# Explicitly define the variables to export, excluding comments
export $(grep -v '^#' "$ENV_FILE" | xargs)

# Run the Go application
go run "$MAIN_FILE"
