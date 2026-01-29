#!/bin/bash
set -euo pipefail

BASE_DIR="/usr/local/bin"
LIB_DIR="$BASE_DIR/lib"
STAGES_DIR="$BASE_DIR/stages"

. "$LIB_DIR/config.sh"
. "$LIB_DIR/logging.sh"

log "Installer started"

bash "$STAGES_DIR/01-disk.sh"
log "Completed stage: 01-disk.sh"
bash "$STAGES_DIR/02-format.sh"
log "Completed stage: 02-format.sh"
bash "$STAGES_DIR/03-mount.sh"
log "Completed stage: 03-mount.sh"


# future
# bash "$STAGES_DIR/02-format.sh"
# log "Completed stage: 02-format.sh"

# bash "$STAGES_DIR/03-mount.sh"
# log "Completed stage: 03-mount.sh"

log "Installer finished"
