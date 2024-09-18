#!/usr/bin/env bash

set -euo pipefail

# Source the USER ID
. .env

# Remove generated files
for F in $(cat .openapi-generator/FILES) ; do
    rm -f "${F}"
done

# Remove test files
for F in $(ls test/*.go) ; do
    rm -f "${F}"
done

# Generate library
docker run --user $USER_ID:$USER_ID --rm --env JAVA_OPTS=-DmaxYamlCodePoints=9999999 -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.6.0 \
    generate \
    --config /local/.openapi-generator/config.yaml \
    --input-spec /local/api/openapi.yaml \
    --output /local \
    --inline-schema-options RESOLVE_INLINE_ENUMS=true \
    --http-user-agent go-netbox/$(cat api/netbox_version)
