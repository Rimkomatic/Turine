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

bash "$STAGES_DIR/031-network.sh"
log "Completed stage: 031-network.sh"

bash "$STAGES_DIR/04-pacstrap.sh"
log "Completed stage: 04-pacstrap.sh"

bash "$STAGES_DIR/05-fstab.sh"
log "Completed stage: 05-fstab.sh"


echo "Completed stage: 05-fstab.sh"

bash "$STAGES_DIR/06-chroot.sh"


# future
# bash "$STAGES_DIR/02-format.sh"
# log "Completed stage: 02-format.sh"

# bash "$STAGES_DIR/03-mount.sh"
# log "Completed stage: 03-mount.sh"

log "Installer finished"
