#!/bin/bash
set -euo pipefail

CONFIG_FILE="/usr/local/bin/lib/config.sh"

. "$CONFIG_FILE"

mapfile -t DISKS < <(lsblk -d -n -o NAME,SIZE,MODEL)

echo "Available disks:"
echo

i=1
for disk in "${DISKS[@]}"; do
  name=$(awk '{print $1}' <<< "$disk")
  size=$(awk '{print $2}' <<< "$disk")
  model=$(cut -d' ' -f3- <<< "$disk")
  echo "[$i] /dev/$name  ($size, $model)"
  ((i++))
done

echo
read -rp "Select disk number: " choice </dev/tty

if ! [[ "$choice" =~ ^[0-9]+$ ]] || (( choice < 1 || choice > ${#DISKS[@]} )); then
  echo "Invalid selection"
  exit 1
fi

DISK="/dev/$(awk '{print $1}' <<< "${DISKS[$((choice - 1))]}")"

echo
echo -e "\033[0;33mWARNING: This will ERASE ALL DATA on $DISK\033[0m"
read -rp "Type the full disk path to confirm: " confirm </dev/tty

[[ "$confirm" == "$DISK" ]] || { echo "Confirmation failed"; exit 1; }

wipefs -a "$DISK"

parted -s "$DISK" mklabel gpt
parted -s "$DISK" mkpart ESP fat32 1MiB 513MiB
parted -s "$DISK" set 1 esp on
parted -s "$DISK" mkpart primary linux-swap 513MiB 8.5GiB
parted -s "$DISK" mkpart primary ext4 8.5GiB 100%

# Notify kernel of partition changes
partprobe "$DISK"
sleep 1

# Detect real partition names (works for SATA, NVMe, etc.)
mapfile -t PARTS < <(lsblk -ln -o NAME "$DISK" | tail -n +2)

EFI_PART="/dev/${PARTS[0]}"
SWAP_PART="/dev/${PARTS[1]}"
ROOT_PART="/dev/${PARTS[2]}"

EFI_SIZE="512M"
SWAP_SIZE="8GiB"
SWAP_TYPE="partition"

cp "$CONFIG_FILE" "$CONFIG_FILE.bak"

sed -i \
  -e "s|^TARGET_DISK=.*|TARGET_DISK=\"$DISK\"|" \
  -e "s|^EFI_PARTITION=.*|EFI_PARTITION=\"$EFI_PART\"|" \
  -e "s|^ROOT_PARTITION=.*|ROOT_PARTITION=\"$ROOT_PART\"|" \
  -e "s|^EFI_SIZE=.*|EFI_SIZE=\"$EFI_SIZE\"|" \
  -e "s|^SWAP_TYPE=.*|SWAP_TYPE=\"$SWAP_TYPE\"|" \
  -e "s|^SWAP_SIZE=.*|SWAP_SIZE=\"$SWAP_SIZE\"|" \
  -e "s|^SWAP_PARTITION=.*|SWAP_PARTITION=\"$SWAP_PART\"|" \
  "$CONFIG_FILE"

echo
echo "Disk layout:"
lsblk "$DISK"
echo

echo "Written installer state:"
echo "TARGET_DISK=$DISK"
echo "EFI_PARTITION=$EFI_PART"
echo "SWAP_PARTITION=$SWAP_PART"
echo "ROOT_PARTITION=$ROOT_PART"
echo "EFI_SIZE=$EFI_SIZE"
echo "SWAP_TYPE=$SWAP_TYPE"
echo "SWAP_SIZE=$SWAP_SIZE"
echo
