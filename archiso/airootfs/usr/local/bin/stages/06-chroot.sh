#!/bin/bash
set -euo pipefail

. /usr/local/bin/lib/config.sh

#--------------------------------------------------
# Ask user credentials (override config)
#--------------------------------------------------
read -rp "Enter username: " USERNAME

while :; do
  read -rsp "Enter password for $USERNAME: " USER_PASSWORD; echo
  read -rsp "Confirm password: " CONFIRM; echo
  [ "$USER_PASSWORD" = "$CONFIRM" ] && break
  echo "Passwords do not match, try again"
done

while :; do
  read -rsp "Enter root password: " ROOT_PASSWORD; echo
  read -rsp "Confirm root password: " CONFIRM; echo
  [ "$ROOT_PASSWORD" = "$CONFIRM" ] && break
  echo "Passwords do not match, try again"
done

export USERNAME USER_PASSWORD ROOT_PASSWORD

#--------------------------------------------------
# Disk setup (assumes partitions already exist)
#--------------------------------------------------
mkfs.fat -F32 "$EFI_PARTITION"
mkfs."$ROOT_FS" "$ROOT_PARTITION"

mount "$ROOT_PARTITION" "$MOUNT_POINT"
mount --mkdir "$EFI_PARTITION" "$MOUNT_POINT$EFI_MOUNT"

ROOT_UUID=$(blkid -s PARTUUID -o value "$ROOT_PARTITION")
export ROOT_UUID

#--------------------------------------------------
# Base install
#--------------------------------------------------
pacstrap -K "$MOUNT_POINT" \
  $BASE_PACKAGES \
  $EXTRA_PACKAGES \
  sudo

genfstab -U "$MOUNT_POINT" >> "$MOUNT_POINT/etc/fstab"

#--------------------------------------------------
# Chroot configuration
#--------------------------------------------------
arch-chroot "$MOUNT_POINT" /bin/bash <<EOF
set -euo pipefail

#--------------------------------------------------
# Hostname
#--------------------------------------------------
echo "$HOSTNAME" > /etc/hostname

cat > /etc/hosts <<HOSTS
127.0.0.1   localhost
::1         localhost
127.0.1.1   $HOSTNAME.localdomain $HOSTNAME
HOSTS

#--------------------------------------------------
# Time & Locale
#--------------------------------------------------
ln -sf "/usr/share/zoneinfo/$TIMEZONE" /etc/localtime
hwclock --systohc

sed -i "s/^#\\(${LOCALE}\\)/\\1/" /etc/locale.gen
locale-gen

echo "LANG=$LANG" > /etc/locale.conf
echo "KEYMAP=$KEYMAP" > /etc/vconsole.conf

#--------------------------------------------------
# Users
#--------------------------------------------------
echo "root:$ROOT_PASSWORD" | chpasswd
useradd -m -G wheel -s "$USER_SHELL" "$USERNAME"
echo "$USERNAME:$USER_PASSWORD" | chpasswd

sed -i 's/^# %wheel ALL=(ALL:ALL) ALL/%wheel ALL=(ALL:ALL) ALL/' /etc/sudoers

#--------------------------------------------------
# Initramfs
#--------------------------------------------------
mkinitcpio -P

#--------------------------------------------------
# Bootloader
#--------------------------------------------------
if [ "$INSTALL_BOOTLOADER" = "yes" ] && [ "$BOOTLOADER" = "systemd-boot" ]; then
  bootctl install

  cat > /boot/loader/loader.conf <<LOADER
default arch
timeout 3
editor no
LOADER

  cat > /boot/loader/entries/arch.conf <<ENTRY
title   Arch Linux
linux   /vmlinuz-linux
initrd  /initramfs-linux.img
options root=PARTUUID=$ROOT_UUID rw
ENTRY
fi

#--------------------------------------------------
# Services
#--------------------------------------------------
if [ "$ENABLE_NETWORK" = "yes" ]; then
  systemctl enable NetworkManager
fi

EOF

#--------------------------------------------------
# Cleanup
#--------------------------------------------------
unset USER_PASSWORD ROOT_PASSWORD

echo "Installation complete."

if [ "$REBOOT_AFTER_INSTALL" = "yes" ]; then
  echo "Rebooting..."
  reboot
fi
