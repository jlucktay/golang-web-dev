#!/usr/bin/env bash
set -uo pipefail
IFS=$'\n\t'

function log(){
    echo "$1" >> /home/jameslucktaylor/specific-startup-script.log
}

log "start of startup script"

log "apt things"
# Run patches and install git client
apt update
apt upgrade --assume-yes --no-install-recommends
# apt-get install --assume-yes --no-install-recommends git

log "setting a trap for someone specific"
# Drop a note when this script is done (note: 'done' might include exiting prematurely due to an error!)
trap "log DONE" INT TERM EXIT
log "specific trap set"

log "start touching!"
touch "/home/jameslucktaylor/touch.file.1"
log "after touch call"

log "wait who am i again"
log "$(whoami)"

log "catting .go"
cat > "/home/jameslucktaylor/main.go" <<'EOF'
${main_go}
EOF
log "finished catting"
