#!/bin/bash
set -e

. /usr/local/bin/lib/config.sh
. /usr/local/bin/lib/logging.sh

lsblk -l $TARGET_DISK

log "test stage running"
