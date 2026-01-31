#!/bin/bash
set -euo pipefail

. /usr/local/bin/lib/config.sh

genfstab -U "$MOUNT_POINT" >> "$MOUNT_POINT/etc/fstab"


echo
echo -e "\033[0;32mFstab generated.\033[0m"
echo
