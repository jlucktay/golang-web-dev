#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'
UserHome="/home/jameslucktaylor"
GoSource="$UserHome/main.go"
ServiceName="revolver"
Binary="$UserHome/$ServiceName"

curl https://raw.githubusercontent.com/jlucktay/golang-web-dev/master/033_aws-scaling/04_hands-on/01_challenge/tf/go/main.go --output $GoSource
sudo systemctl stop $ServiceName
go build -o $Binary -a -ldflags '-extldflags "-static"' -v -work $GoSource
sudo systemctl start $ServiceName
