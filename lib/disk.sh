wipefs -a /dev/sda
parted /dev/sda --script mklabel gpt

parted /dev/sda --script mkpart ESP fat32 1MiB 513MiB
parted /dev/sda --script set 1 esp on

parted /dev/sda --script mkpart primary linux-swap 513MiB 8.5GiB
parted /dev/sda --script mkpart primary ext4 8.5GiB 100%
