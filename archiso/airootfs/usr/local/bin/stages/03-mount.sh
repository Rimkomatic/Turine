#!/bin/bash
set -euo pipefail

. /usr/local/bin/lib/config.sh

# Mount root filesystem
mount "$ROOT_PARTITION" "$MOUNT_POINT"

# Mount EFI
mkdir -p "$MOUNT_POINT/boot"
mount "$EFI_PARTITION" "$MOUNT_POINT/boot"

echo
echo -e "\033[0;32mMounting complete.\033[0m"
echo

echo "Mounted filesystems:"
lsblk "$TARGET_DISK"
echo
