#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

gcloud --project golang-web-dev-227919 compute instances list --format=json | jq -r '.[] | { n: .name, c: .creationTimestamp, z: (.zone | split("/"))[-1] }'
