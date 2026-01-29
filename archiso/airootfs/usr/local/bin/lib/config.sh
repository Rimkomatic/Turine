#!/bin/bash

# turine installer – configuration
# sourced by all stages

#--------------------------------------------------
# Installer context
#--------------------------------------------------

INSTALL_NAME="turine"
MOUNT_POINT="/mnt"

#--------------------------------------------------
# Disk layout (filled by 01-disk.sh)
#--------------------------------------------------

TARGET_DISK=""
EFI_PARTITION=""
ROOT_PARTITION=""

WIPE_DISK="yes"            # yes | no
BOOT_MODE="uefi"           # uefi | bios

EFI_SIZE="512M"
ROOT_FS="ext4"             # ext4 | btrfs | xfs

#--------------------------------------------------
# Mount points
#--------------------------------------------------

EFI_MOUNT="/boot"
ROOT_MOUNT="/"
#--------------------------------------------------
# Swap
#--------------------------------------------------

SWAP_TYPE="file"      # file | partition | none
SWAP_SIZE="8GiB"
SWAP_PARTITION=""

#--------------------------------------------------
# System identity
#--------------------------------------------------

HOSTNAME="turine"
TIMEZONE="Asia/Kolkata"

LOCALE="en_US.UTF-8"
LANG="en_US.UTF-8"
KEYMAP="us"

#--------------------------------------------------
# Users
#--------------------------------------------------

ROOT_PASSWORD="changeme"

USERNAME="rik"
USER_PASSWORD="changeme"
USER_SHELL="/bin/bash"

#--------------------------------------------------
# Packages
#--------------------------------------------------

BASE_PACKAGES="base linux linux-firmware"
EXTRA_PACKAGES="vim git networkmanager"

#--------------------------------------------------
# Bootloader
#--------------------------------------------------

INSTALL_BOOTLOADER="yes"
BOOTLOADER="systemd-boot"   # systemd-boot | grub

#--------------------------------------------------
# Behavior flags
#--------------------------------------------------

ENABLE_NETWORK="yes"
REBOOT_AFTER_INSTALL="yes"
