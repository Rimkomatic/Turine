# to check the partitions

fdisk -l

fdisk /dev/name

# partition
n        
<enter>  
<enter>  
+512M    
t
<enter>  
1        


n
<enter>
<enter>
+8G        
t
<enter>
19         


n
<enter>
<enter>
<enter>  
t
<enter>
23       


w

mkfs.fat -F32 /dev/sda1
mkswap /dev/sda2
mkfs.ext4 /dev/sda3


swapon /dev/sda2



# Mounting

mount /dev/sda3 /mnt

mkdir /mnt/boot
mount /dev/sda1 /mnt/boot




pacman -Syy
pacman -S reflector
cp /etc/pacman.d/mirrorlist /etc/pacman.d/mirrorlist.bak


reflector -c "IN" -f 12 -l 10 -n 12 --save /etc/pacman.d/mirrorlist







pacstrap /mnt base linux linux-firmware vim nano

genfstab -U /mnt >> /mnt/etc/fstab
arch-chroot /mnt
passwd


pacman -S networkmanager
systemctl enable NetworkManager







useradd -m -G wheel -s /bin/bash rik
passwd rik

pacman -S sudo

#edit this line
EDITOR=nano visudo
#uncomment this line
%wheel ALL=(ALL:ALL) ALL



systemctl enable systemd-timesyncd


exit
umount -R /mnt
reboot




