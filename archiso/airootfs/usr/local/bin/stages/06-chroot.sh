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
  [[ "$USER_PASSWORD" == "$CONFIRM" ]] && break
  echo "Passwords do not match, try again"
done

while :; do
  read -rsp "Enter root password: " ROOT_PASSWORD; echo
  read -rsp "Confirm root password: " CONFIRM; echo
  [[ "$ROOT_PASSWORD" == "$CONFIRM" ]] && break
  echo "Passwords do not match, try again"
done

export USERNAME USER_PASSWORD ROOT_PASSWORD

#--------------------------------------------------
# Derive ROOT_UUID automatically
#--------------------------------------------------
ROOT_UUID="$(blkid -s PARTUUID -o value "$ROOT_PARTITION")"


install -Dm755 \
  /usr/local/bin/turine-setup/turine \
  "$MOUNT_POINT/usr/local/bin/turine-setup/turine"

#--------------------------------------------------
# Enter chroot
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

sed -i "s/^#\(${LOCALE}\)/\1/" /etc/locale.gen || true
locale-gen

echo "LANG=$LANG" > /etc/locale.conf
echo "KEYMAP=$KEYMAP" > /etc/vconsole.conf

#--------------------------------------------------
# Packages
#--------------------------------------------------
pacman -Sy --noconfirm $EXTRA_PACKAGES

#--------------------------------------------------
# Users
#--------------------------------------------------
echo "root:$ROOT_PASSWORD" | chpasswd

useradd -m -G wheel -s "$USER_SHELL" "$USERNAME"
echo "$USERNAME:$USER_PASSWORD" | chpasswd

sed -i 's/^# %wheel ALL=(ALL:ALL) ALL/%wheel ALL=(ALL:ALL) ALL/' /etc/sudoers

#--------------------------------------------------
# Enable passwordless sudo (required)
#--------------------------------------------------
echo "%wheel ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/99-wheel-nopasswd
chmod 440 /etc/sudoers.d/99-wheel-nopasswd

#--------------------------------------------------
# Initramfs
#--------------------------------------------------
mkinitcpio -P

#--------------------------------------------------
# Bootloader (systemd-boot)
#--------------------------------------------------
if [[ "$INSTALL_BOOTLOADER" == "yes" && "$BOOTLOADER" == "systemd-boot" ]]; then
  bootctl install

  cat > /boot/loader/loader.conf <<LOADER
default arch
timeout 3
editor no
LOADER

  cat > /boot/loader/entries/arch.conf <<ENTRY
title   arch
linux   /vmlinuz-linux
initrd  /initramfs-linux.img
options root=PARTUUID=$ROOT_UUID rw
ENTRY
fi

#--------------------------------------------------
# System services
#--------------------------------------------------
if [[ "$ENABLE_NETWORK" == "yes" ]]; then
  systemctl enable NetworkManager
fi

#--------------------------------------------------
# Turine first-boot handoff (run with sudo)
#--------------------------------------------------
sudo -u "$USERNAME" </dev/tty >/dev/tty 2>&1 /usr/local/bin/turine-setup/turine

EOF
