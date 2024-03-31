#!/bin/bash

# This script generates boilerplate code from the OpenAPI spec
# It uses oapi-codegen to generate Go code from the OpenAPI spec

# This script requires the following tools to be installed:
# - oapi-codegen

# Generate the code from the OpenAPI specification
echo "Generating code from OpenAPI spec..."

# -------------------------------------------------------------------------------------------------

echo "Generating code for the schemas..."

echo "Generating code for the schema: auth..."
oapi-codegen -config api/code-gen-configs/schemas/auth.yaml api/openapi/schemas/auth.yaml

echo "Generating code for the schema: errors..."
oapi-codegen -config api/code-gen-configs/schemas/errors.yaml api/openapi/schemas/errors.yaml

echo "Generating code for the schema: project..."
oapi-codegen -config api/code-gen-configs/schemas/project.yaml api/openapi/schemas/project.yaml

echo "Generating code for the schema: task..."
oapi-codegen -config api/code-gen-configs/schemas/task.yaml api/openapi/schemas/task.yaml

echo "Generating code for the schema: user..."
oapi-codegen -config api/code-gen-configs/schemas/user.yaml api/openapi/schemas/user.yaml

# -------------------------------------------------------------------------------------------------

echo "Generating code for the servers..."

echo "Generating code for the server: Auth Service..."
oapi-codegen -config api/code-gen-configs/server/auth.yaml api/openapi/server/auth.yaml

echo "Generating code for the server: Project Service..."
oapi-codegen -config api/code-gen-configs/server/project.yaml api/openapi/server/project.yaml

echo "Generating code for the server: User Service..."
oapi-codegen -config api/code-gen-configs/server/user.yaml api/openapi/server/user.yaml

# -------------------------------------------------------------------------------------------------
