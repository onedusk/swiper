# Page 13

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.2.2 Configure Hyper-V for Bridge Network

• Open Hyper-V Manager.
• From the Actions pane, select Virtual Switch Manager.
• Choose the type of virtual switch as External, then select Create Virtual Switch.
• Enter a name for the virtual switch as Bridge.
• Choose the wifi network adapter (NIC) that you want to use, then select OK.
You’ll be prompted with a warning that the change may disrupt your network connectivity; select Yes if you’re happy to
continue.
Create .wslconfig file on C:/Users/user-name/ and add the following lines to connect to configured virtual Bridge Network.
[wsl2]
networkingMode = bridged
vmSwitch = Bridge
ipv6 = true

1.2. 2. Developing with the SDK

9


```

## Images

![Image from page 13](images/page_13_img_001.ppm)

![Image from page 13](images/page_13_img_002.ppm)

