#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'
UserHome="/home/jameslucktaylor"
LogFile="$UserHome/gce.startup.log"
GoSource="$UserHome/main.go"
ServiceName="revolver"
Binary="$UserHome/$ServiceName"
ServiceFile="/etc/systemd/system/$ServiceName.service"

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

log "Building '$ServiceName' binary..."
go build -o $Binary -a -ldflags '-extldflags "-static"' -v -work $GoSource >> $LogFile 2>&1
log "Built '$ServiceName' binary."

log "Catting '$ServiceFile'..."
tee $ServiceFile <<'EOF'
[Unit]
Description=Go Server

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
systemctl enable $ServiceName >> $LogFile 2>&1

log "Activate the service..."
systemctl start $ServiceName >> $LogFile 2>&1

log "Check if systemd started it..."
systemctl status $ServiceName >> $LogFile 2>&1
