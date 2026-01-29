#!/bin/sh

LOG_FILE="/tmp/turine-install.log"

log() {
    printf '[INFO] %s\n' "$*" | tee -a "$LOG_FILE"
}

warn() {
    printf '[WARN] %s\n' "$*" | tee -a "$LOG_FILE" >&2
}

die() {
    printf '[FATAL] %s\n' "$*" | tee -a "$LOG_FILE" >&2
    exit 1
}

# aliases (optional sugar)
log_info()    { log "$@"; }
log_warn()    { warn "$@"; }
log_error()   { warn "$@"; }
log_success() { log "$@"; }
