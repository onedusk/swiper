# Page 20

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.3.1 Test Setup (CHIP Tool)
A host-based chip-tool can be used as a commissioner to commission and control a Matter device. During the previous
install.sh step, the chip-tool is generated under the folder:
${ESP_MATTER_PATH}/connectedhomeip/connectedhomeip/out/host

Note: macOS Users: To use chip-tool with BLE commissioning on macOS, you must install the Bluetooth Central
Matter Client Developer Mode Profile. It enables Matter commissioning, and may require periodic re-installation.
Instructions to download the profile can be found in the profile installation section

2.3.1.1 Commissioning
Use chip-tool in interactive mode to commission the device:
chip-tool interactive start

In the above commands:
• 0x7283 is the randomly chosen node_id
• 20202021 is the setup_passcode
• 3840 is the discriminator
Above method commissions the device using setup passcode and discriminator. Device can also be commissioned using
manual pairing code or QR code.
To Commission the device using manual pairing code 34970112332
Above default manual pairing code contains following values:
Version:
Custom flow:
Discriminator:
Passcode:

0
0
(STANDARD)
3840
20202021

To commission the device using QR code MT:Y.K9042C00KA0648G00
Above QR Code contains the below default values:
Version:
Vendor ID:
ProductID:
Custom flow:
Discovery Bitmask:
Long discriminator:
Passcode:

0
65521
(0xFFF1)
32768
(0x8000)
0
(STANDARD)
0x02
(BLE)
3840
(0xf00)
20202021

Alternatively, you can scan the below QR code image using Matter commissioners.

If QR code is not visible, paste the below link into the browser and scan the QR code.

16

Chapter 1. Table of Contents


```

## Images

![Image from page 20](images/page_20_img_001.ppm)

![Image from page 20](images/page_20_img_002.ppm)

