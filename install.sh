#!/usr/bin/bash
set -euo pipefail

LOG="/tmp/turine-install.log"

echo "==================================" | tee -a "$LOG"
echo "Turine installer started"          | tee -a "$LOG"
echo "Date: $(date)"                    | tee -a "$LOG"
echo "Running inside ArchISO"           | tee -a "$LOG"
echo "==================================" | tee -a "$LOG"

sleep 3

echo "Installer exiting successfully."  | tee -a "$LOG"
