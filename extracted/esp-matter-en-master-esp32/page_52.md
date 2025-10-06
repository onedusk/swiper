# Page 52

## Text Content

```
ESP-Matter Programming Guide, Release latest

4.1.3 Certification Declaration (CD)
A Certification Declaration (CD) is a cryptographic document that allows a Matter device to assert its protocol compliance. Once your product is certified, the CSA creates a CD for that device. The CD should then be included in the
device firmware by the device manufacturer.
4.1.4 Setup Passcode, Discriminator and Onboarding Payload
The unique setup passcode serves as the proof of possession and is also used to compute the shared secret during
commissioning. The corresponding SPAKE2+ verifier of the passcode is installed on the device and not the actual
passcode.
The discriminator is used to easily distinguish between devices to provide a seamless experience during commissioning.
The onboarding payload is the QR code and the manual pairing code that assists a commissioner (like a phone app) to
allow onboarding a device into the Matter network. The QR code and/or the manual pairing code are generally printed
on the packaging of the device.
4.1.5 Manufacturing Partition
Espressif’s SDK for Matter uses a separate manufacturing partition to store all the information mentioned above. Because
the DACs are unique to every device, the manufacturing partition will also be unique per device. Thus by moving all the
typical per device unique fields into the manufacturing partition, the rest of the components like the bootloader, firmware
image are common across all your devices. You can refer the Manufacturing section below for creating a large number
of manufacturing partition images.
Your manufacturing line needs to ensure that these unique manufacturing partition images are correctly written to each
device and the appropriate QR code images associated with each device. You may also opt for Espressif’s pre-provisioning
service that pre-provisions these unique images before shipping the modules and provides a manifest (CSV file) along with
QR code images bundle.

1.4.2 4.2 Over-the-Air (OTA) Updates
Matter devices must support OTA firmware updates, either by using Matter-based OTA or vendor specific means.
In case of Matter OTA, there’s an OTA provider that assists an OTA requestor to get upgraded. The SDK examples support
Matter OTA requestor role out of the box. The OTA provider could be a manufacturer specific phone app or any Matter
node that has internet connectivity.
Alternatively, ESP RainMaker OTA service can also be used to upgrade the firmware on the devices remotely. As opposed
to the Matter OTA, ESP RainMaker OTA allows you the flexibility of delivering the OTA upgrades incrementally or to
groups of devices.

1.4.3 4.3 Manufacturing
4.3.1 Mass Manufacturing Utility
For commissioning a device into the Matter Fabric, the device requires the following information:
• Device Attestation Certificate (DAC) and Certification Declaration (CD): verified by commissioner to determine whether a device is a Matter certified product or not.
• Discriminator: advertised during commissioning to easily distinguish between advertising devices.

48

Chapter 1. Table of Contents


```

