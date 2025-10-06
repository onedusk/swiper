# Page 15

## Text Content

```
ESP-Matter Programming Guide, Release latest

The new kernel image will be built.
Copy the new kernel
cp arch/x86/boot/bzImage /mnt/path/to/kernel/bluetooth-bzImage

1.3.2 Configure WSL to use new custom kernel image
Add the following line to the created .wslconfig file.
[wsl2]
kernel=c:\\users\\<user>\\bluetooth-bzImage

Replace the path with the path of new custom kernel built.
1.3.3 Attaching Bluetooth module to WSL2
Get the BUSID of the bluetooth module. [Tested using usbipd-win 4.0.0]
usbipd list

Attach the bluetooth module to WSL2 instance using usbipd.
usbipd attach --wsl --busid=<BUSID>
$ lsusb
Bus 002 Device 001: ID 1d6b:0003 Linux Foundation 3.0 root hub
Bus 001 Device 003: ID 8087:0029 Intel Corp. AX201 Bluetooth
Bus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub

The bluetooth module should be available to WSL.
1.3.4 Testing Bluetooth
Install bluez library and scan for bluetooth devices.
sudo apt install bluez

Start scanning for available Bluetooth devices.
bluetoothctl scan on

The bluetooth discovery should start.

1.2. 2. Developing with the SDK

11


```

## Images

![Image from page 15](images/page_15_img_001.ppm)

![Image from page 15](images/page_15_img_002.ppm)

