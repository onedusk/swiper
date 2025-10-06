# Page 14

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.2.2.1 Checking the configuration
ifconfig

The IPv4 Address assigned to the eth0 will be from the local network and IPv6 Address will also be assigned to it.
avahi-browse _matter._tcp -r

If the configuration was done correctly, all the matter operational nodes performing mDns advertisement on the local
network will be shown.
1.2.2.2 Troubleshoot
If no IP address assigned to WSL or same IP address as host, uspipd fails with tcp connect error, not able to ping
external network.
1) Run wsl –shutdown and wait for WSL instance to close.
2) Run Hyper-V Manager and open Virtual Switch Manager from Actions Pane.
3) Disconnect and Reconnect to Wi-Fi.
4) Start WSL instance.
5) If still no IP is assigned, run sudo dhclient
1.3 Providing access to bluetooth
Bluetooth support is missing in default WSL kernel. To add support for bluetooth, WSL kernel needs to be recompiled
with right drivers.
1.3.1 Building custom kernel for bluetooth access in WSL2
git clone --depth 1 --branch linux-msft-wsl-6.1.21.2 https://github.com/microsoft/
,→WSL2-Linux-Kernel.git

Replace branch with the latest available WSL-Linux-Kernel tag.
cd WSL2-Linux-Kernel
git checkout linux-msft-wsl-6.1.21.2
cp /proc/config.gz config.gz
gunzip config.gz
mv config .config
sudo make menuconfig

Select the features to be enabled in the kernel:
1) Enable Networking support ->Bluetooth subsystem support.
2) Enable Networking Support ->Bluetooth Subsystem Support ->Bluetooth device drivers -> HCI USB driver.
3) Save the config file.
sudo make -j$(getconf _NPROCESSORS_ONLN) && sudo make modules_install -j$(getconf _
,→NPROCESSORS_ONLN) && sudo make install -j$(getconf _NPROCESSORS_ONLN)

10

Chapter 1. Table of Contents


```

