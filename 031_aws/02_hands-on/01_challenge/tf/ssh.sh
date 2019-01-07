#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

gcloud --project golang-web-dev-227919 compute ssh --zone europe-west2-a hello-go
