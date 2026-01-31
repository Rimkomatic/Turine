#!/bin/bash
set -euo pipefail

. /usr/local/bin/lib/config.sh


pacstrap "$MOUNT_POINT" base linux linux-firmware

echo
echo -e "\033[0;32mBase Packages Installed\033[0m"
echo
