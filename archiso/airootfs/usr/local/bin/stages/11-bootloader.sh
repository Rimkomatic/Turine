bootctl install

blkid /dev/sda3

cat <<EOF > /boot/loader/loader.conf
default arch
timeout 3
editor no
EOF

cat <<EOF > /boot/loader/entries/arch.conf
title Arch Linux
linux /vmlinuz-linux
initrd /initramfs-linux.img
options root=UUID=XXXX rw
EOF
