#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

for instance in $(./ssh.list.sh | jq -r '.n'); do
    gcloud --project golang-web-dev-227919 compute scp --zone europe-west2-a gce.bounce.sh "${instance}:~"
done
