#!/bin/bash
set -euo pipefail

. /usr/local/bin/lib/config.sh

# --------------------------------------------------
# Ask connection type
# --------------------------------------------------
echo
echo "Select network type:"
echo "1) Wi-Fi"
echo "2) Ethernet"
read -rp "Choice [1-2]: " NET_TYPE
echo

# --------------------------------------------------
# Wi-Fi path (UNCHANGED)
# --------------------------------------------------
if [ "$NET_TYPE" = "1" ]; then
  echo "Starting iwd..."
  systemctl start iwd
  sleep 5

  # --------------------------------------------------
  # Detect Wi-Fi device
  # --------------------------------------------------
  WIFI_DEV=$(iw dev | awk '$1=="Interface"{print $2; exit}')

  if [ -z "$WIFI_DEV" ]; then
    echo "ERROR: No Wi-Fi device found."
    exit 1
  fi

  echo "Using Wi-Fi device: $WIFI_DEV"

  # Ensure powered
  iwctl device "$WIFI_DEV" set-property Powered on

  # --------------------------------------------------
  # Scan + show networks
  # --------------------------------------------------
  echo
  echo "Scanning Wi-Fi networks..."
  iwctl station "$WIFI_DEV" scan
  sleep 5
  iwctl station "$WIFI_DEV" get-networks

  # --------------------------------------------------
  # Ask user
  # --------------------------------------------------
  echo
  read -rp "Enter Wi-Fi SSID: " SSID
  echo

  # --------------------------------------------------
  # Connect (iwd prompts internally)
  # --------------------------------------------------
  iwctl station "$WIFI_DEV" connect "$SSID"
fi

# --------------------------------------------------
# Ethernet path
# --------------------------------------------------
if [ "$NET_TYPE" = "2" ]; then
  ETH_DEV=$(ip -o link show | awk -F': ' '!/lo|wl/{print $2; exit}')

  if [ -z "$ETH_DEV" ]; then
    echo "ERROR: No Ethernet device found."
    exit 1
  fi

  echo "Using Ethernet device: $ETH_DEV"
  ip link set "$ETH_DEV" up
fi

# --------------------------------------------------
# Bring up networking (installer-safe)
# --------------------------------------------------
systemctl enable --now systemd-networkd systemd-resolved
ln -sf /run/systemd/resolve/stub-resolv.conf /etc/resolv.conf

# --------------------------------------------------
# Wait for DHCP
# --------------------------------------------------
echo "Waiting for DHCP..."
for _ in {1..15}; do
  ip route | grep -q default && break
  sleep 1
done

# --------------------------------------------------
# Wait for DNS
# --------------------------------------------------
echo "Waiting for DNS..."
for _ in {1..15}; do
  resolvectl status >/dev/null 2>&1 && break
  sleep 1
done

# --------------------------------------------------
# Hard network gate
# --------------------------------------------------
echo "Checking network connectivity..."
ping -c 1 archlinux.org >/dev/null || {
  echo "ERROR: Network not available."
  exit 1
}

echo "Network connected successfully."
