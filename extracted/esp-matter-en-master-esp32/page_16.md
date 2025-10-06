# Page 16

## Text Content

```
ESP-Matter Programming Guide, Release latest

Tested Bluetooth modules:
• Intel AX201 series and above
• Marvell AVASTAR Bluetooth Radio Adapter.
1.4 Final .wslconfig file
[wsl2]
kernel = D:\\custom-kernel\\bluetooth-bzImage
networkingMode = bridged
vmSwitch = Bridge
ipv6 = true

Replace the kernel path appropriately.
2.1.2 Getting the Repository
The Prerequisites for ESP-IDF:
• Please get the Prerequisites for ESP-IDF. For beginners, please check step by step installation guide for esp-idf.
Note: git clone command accepts the optional argument --jobs N, which can significantly speed up the process
by parallelizing submodule cloning. Consider using this option when cloning repositories.

2.1.3 Configuring the Environment
This should be done each time a new terminal is opened
cd esp-idf; source ./export.sh; cd ..

1.2.2 2.2 ESP Matter Setup
There are two options to setup esp-matter, you can select one according to demand:
• ESP matter repository, including esp-matter SDK and tools (e.g., CHIP-tool, CHIP-cert, ZAP, …).
• ESP matter component, including esp-matter SDK.
2.2.1 ESP-Matter Repository
2.2.1.1 Getting the Repository
The Prerequisites for Matter:
• Please get the Prerequisites for Matter.
Cloning the esp-matter repository takes a while due to a lot of submodules in the upstream connectedhomeip, so if you
want to do a shallow clone use the following command:
• For Linux host:

12

Chapter 1. Table of Contents


```

