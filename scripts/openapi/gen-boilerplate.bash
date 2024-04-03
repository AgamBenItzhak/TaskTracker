#!/bin/bash

# This script generates boilerplate code from the OpenAPI spec
# It uses oapi-codegen to generate Go code from the OpenAPI spec

# This script requires the following tools to be installed:
# - oapi-codegen

# Check if oapi-codegen is installed
if ! command -v oapi-codegen &> /dev/null; then
    echo "oapi-codegen is not installed. Please install oapi-codegen before running this script."
    exit 1
fi

# Generate the code from the OpenAPI specification
echo "Generating code from OpenAPI spec..."

# -------------------------------------------------------------------------------------------------

echo "Generating code for the schemas..."

echo "Generating code for the schema: auth..."
oapi-codegen -config api/code-gen-configs/schemas/auth.yaml api/openapi/schemas/auth.yaml || { echo "Failed to generate code for the schema: auth"; exit 1; }

echo "Generating code for the schema: errors..."
oapi-codegen -config api/code-gen-configs/schemas/errors.yaml api/openapi/schemas/errors.yaml || { echo "Failed to generate code for the schema: errors"; exit 1; }

echo "Generating code for the schema: project..."
oapi-codegen -config api/code-gen-configs/schemas/project.yaml api/openapi/schemas/project.yaml || { echo "Failed to generate code for the schema: project"; exit 1; }

echo "Generating code for the schema: task..."
oapi-codegen -config api/code-gen-configs/schemas/task.yaml api/openapi/schemas/task.yaml || { echo "Failed to generate code for the schema: task"; exit 1; }

echo "Generating code for the schema: member..."
oapi-codegen -config api/code-gen-configs/schemas/member.yaml api/openapi/schemas/member.yaml || { echo "Failed to generate code for the schema: member"; exit 1; }

# -------------------------------------------------------------------------------------------------

echo "Generating code for the servers..."

echo "Generating code for the server: Auth Service..."
oapi-codegen -config api/code-gen-configs/server/auth.yaml api/openapi/server/auth.yaml || { echo "Failed to generate code for the server: Auth Service"; exit 1; }

echo "Generating code for the server: Project Service..."
oapi-codegen -config api/code-gen-configs/server/project.yaml api/openapi/server/project.yaml || { echo "Failed to generate code for the server: Project Service"; exit 1; }

echo "Generating code for the server: member Service..."
oapi-codegen -config api/code-gen-configs/server/member.yaml api/openapi/server/member.yaml || { echo "Failed to generate code for the server: member Service"; exit 1; }

# -------------------------------------------------------------------------------------------------

echo "Generation completed successfully."