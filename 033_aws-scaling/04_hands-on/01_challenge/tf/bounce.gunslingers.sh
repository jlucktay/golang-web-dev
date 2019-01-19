#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

for instance in $(./ssh.list.sh | jq -r '.n'); do
    gcloud --project golang-web-dev-227919 compute ssh --zone europe-west2-a "$instance" -- ./gce.bounce.sh
done
