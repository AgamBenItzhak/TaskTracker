#!/bin/bash

# This script generates Go code from the SQL schema

# This script requires the following tools to be installed:
# - sqlc

# Generate the code from the SQL schema
echo "Generating code from SQL schema..."
cd database

# Check if sqlc is installed
if ! command -v sqlc &> /dev/null; then
    echo "sqlc is not installed. Please install sqlc before running this script."
    exit 1
fi

# Generate the code
if ! sqlc generate; then
    echo "Failed to generate code from SQL schema."
    exit 1
fi

echo "Code generated successfully."