# Page 11

## Text Content

```
ESP-Matter Programming Guide, Release latest

• Here onwards process for setting esp-matter and building examples is same as other hosts.
• Please clone the repositories from inside the WSL environment and not inside a mounted directory.
For using CHIP tool on WSL, please check Using CHIP-tool in WSL.
For using VSCode for development, please check Developing in WSL.
1. Working with the CHIP Tool in WSL2
The CHIP Tool (chip-tool) is a Matter controller implementation that allows to commission a Matter device into the
network and to communicate with it using Matter messages, which may encode Data Model actions, such as cluster
commands.
The tool also provides other utilities specific to Matter, such as parsing of the setup payload or performing discovery
actions.
The CHIP Tool requires access to the local network and bluetooth.
1.1 Requirements
• Windows 11 64-Bit Pro/Enterprise/Education [for Hyper-V Manager]
1.2 Providing access to local network
WSL2 does not share an IP address with your computer. Because WSL2 was implemented with Hyper-V, it runs with a
virtualized ethernet adapter. Your computer hides WSL2 behind a NAT where WSL2 has its own unique IP.
To provide an IP address from the local network, WSL2 instance needs to be connected to a virtual Bridge through
Hyper-V Manager.
1.2.1 Enabling Hyper-V for use on Windows 10
1.2.1.1 Enable Hardware Virtualization in BIOS
• In the Startup Menu, enter the BIOS setup.
• In the BIOS Setup Utility,open the Configuration or Security tab.
• Enable the Virtualization Technology option
1.2.1.2 Setting Up Hyper-V
Ensure that hardware virtualization support is turned on in the BIOS settings Save the BIOS settings and boot up the
machine normally.
• Right click on the Windows button and select ‘Apps and Features’.
• Select Programs and Features on the right under related settings.
• Select Turn Windows Features on or off.
• Select Hyper-V and click OK.

1.2. 2. Developing with the SDK

7


```

