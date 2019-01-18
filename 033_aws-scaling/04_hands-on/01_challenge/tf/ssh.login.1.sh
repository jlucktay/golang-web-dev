#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

if [[ -z "$1" ]]; then
    gcloud --project golang-web-dev-227919 compute ssh --zone europe-west2-a "$(gcloud --project golang-web-dev-227919 compute instances list --format=json | jq -r '.[].name' | head -n 1)"
else
    gcloud --project golang-web-dev-227919 compute ssh --zone europe-west2-a "$1"
fi
