#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function log(){
    echo "[$(date '+%Y%m%d.%H%M%S.%N%z')] $1" >> /home/jameslucktaylor/gce.startup.log
}

# Timestamp start
log "cloud-init: start"

# Drop a note when this script is done (note: 'done' might include exiting prematurely due to an error!)
trap "log 'cloud-init: finish'" INT TERM EXIT

# log "Catting '.toprc'..."
# sudo -u jameslucktaylor tee /home/jameslucktaylor/.toprc <<'EOF'
# ${toprc}
# EOF
# log "Catted '.toprc'."

log "Running 'apt'..."
# Run patches and install golang
log "'apt update'..."
apt update
log "'apt upgrade'..."
apt upgrade --assume-yes --no-install-recommends
log "'apt install'..."
apt install gcc build-essential golang-go --assume-yes --no-install-recommends
log "'apt autoremove'..."
apt autoremove --assume-yes
log "Finished 'apt'."

log "Fetching main.go from GitHub..."
curl https://raw.githubusercontent.com/jlucktay/golang-web-dev/master/033_aws-scaling/04_hands-on/01_challenge/tf/go/main.go | sudo -u jameslucktaylor tee /home/jameslucktaylor/main.go
log "Fetched main.go from GitHub."

log "Building 'revolver' binary..."
go build -o revolver -a -ldflags '-extldflags "-static"' main.go
# go build -o revolver -a -installsuffix cgo -ldflags '-extldflags "-static"' main.go
log "Built 'revolver' binary."

# # echo "Hello, World" > index.html
# # nohup busybox httpd -f -p 8080 &
