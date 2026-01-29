#!/bin/bash
set -euo pipefail

. /usr/local/bin/lib/config.sh

# Format EFI
mkfs.fat -F32 "$EFI_PARTITION"

# Format and enable swap
if [[ "$SWAP_TYPE" == "partition" ]]; then
  mkswap "$SWAP_PARTITION"
  swapon "$SWAP_PARTITION"
fi

# Format root
mkfs.ext4 "$ROOT_PARTITION"

echo
echo -e "\033[0;32mFormatting complete using installer state.\033[0m"
echo

echo "Current disk structure:"
lsblk "$TARGET_DISK"
echo
