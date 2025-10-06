# Page 55

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.5 5. Security Considerations
1.5.1 5.1 Overview
This guide provides an overview of the overall security features that should be considered while designing the products
with Matter framework on ESP32 SoCs.
High level security goals are as follows:
1. Preventing untrustworthy code from being executed
2. Securing device identity (e.g., Matter DAC Private Key)
3. Secure storage for confidential data

1.5.2 5.2 Platform Security
5.2.1 Secure Boot
The Secure Boot feature ensures that only authenticated software can execute on the device. The Secure Boot process
forms a chain of trust by verifying all mutable software entities involved in the boot-up process. Signature verification
happens during both boot-up as well as in OTA updates.
Please refer to Secure Boot V2 guide for detailed documentation about this feature in ESP32.
5.2.2 Flash Encryption
The Flash Encryption feature helps to encrypt the contents on the off-chip flash memory and thus provides the confidentiality aspect to the software or data stored in the flash memory.
Please refer to Flash Encryption guide for detailed documentation about this feature in ESP32.

1.5.3 5.3 Product Security
5.3.1 Secure Storage
Secure storage refers to the application-specific data that can be stored in a secure manner on the device, i.e., off-chip flash
memory. This is typically a read-write flash partition and holds device specific configuration data, e.g., Wi-Fi credentials.
ESP-IDF provides the NVS (Non-volatile Storage) management component which allows encrypted data partitions.
This feature is tied with the platform flash encryption feature described earlier.
Please refer to the NVS Encryption for detailed documentation on the working and instructions to enable this feature in
ESP32.

1.5. 5. Security Considerations

51


```

