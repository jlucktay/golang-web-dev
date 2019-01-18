#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'
UserHome="/home/jameslucktaylor"
LogFile="$UserHome/gce.startup.log"
GoSource="$UserHome/main.go"
Binary="$UserHome/revolver"
ServiceFile="/etc/systemd/system/revolver.service"

function log(){
    echo "[$(date '+%Y%m%d.%H%M%S.%N%z')] $1" | sudo -u jameslucktaylor tee --append $LogFile
}

# Timestamp start
log "cloud-init: start"

# Drop a note when this script is done (note: 'done' might include exiting prematurely due to an error!)
trap "log 'cloud-init: finish'" INT TERM EXIT

# log "Catting '.toprc'..."
# sudo -u jameslucktaylor tee /home/jameslucktaylor/.toprc <<'EOF'
# $
# {
# toprc
# }
# EOF
# log "Catted '.toprc'."

# Run patches and install Go, GCC, etc
log "Running 'apt'..."
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
curl https://raw.githubusercontent.com/jlucktay/golang-web-dev/master/033_aws-scaling/04_hands-on/01_challenge/tf/go/main.go | sudo -u jameslucktaylor tee $GoSource
log "Fetched main.go from GitHub."

log "Building 'revolver' binary..."
go build -o $Binary -a -ldflags '-extldflags "-static"' -v -work $GoSource >> $LogFile 2>&1
log "Built 'revolver' binary."

log "Catting '$ServiceFile'..."
tee $ServiceFile <<'EOF'
[Unit]
Description=Revolver - Go Server

[Service]
EOF

echo "ExecStart=$Binary" >> $ServiceFile
echo "WorkingDirectory=$UserHome" >> $ServiceFile

tee --append $ServiceFile <<'EOF'
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
EOF
log "Catted '$ServiceFile'."

log "Add the service to systemd..."
systemctl enable $ServiceFile >> $LogFile

log "Activate the service..."
systemctl start $ServiceFile >> $LogFile

log "Check if systemd started it..."
systemctl status $ServiceFile >> $LogFile
