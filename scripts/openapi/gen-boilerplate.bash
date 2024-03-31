#!/bin/bash

# This script generates boilerplate code from the OpenAPI spec
# It uses oapi-codegen to generate Go code from the OpenAPI spec

# This script requires the following tools to be installed:
# - oapi-codegen

# Generate the code from the OpenAPI specification
echo "Generating code from OpenAPI spec..."

echo "Generating code for the schema..."
oapi-codegen -config api/code-gen-configs/schema.yaml api/openapi/schema.yaml

echo "Generating code for the server..."

echo "Generating code for the server: User Service..."
oapi-codegen -config api/code-gen-configs/server/user.yaml api/openapi/server/user.yaml

echo "Generating code for the server: Project Service..."
oapi-codegen -config api/code-gen-configs/server/project.yaml api/openapi/server/project.yaml

echo "files generated successfully"