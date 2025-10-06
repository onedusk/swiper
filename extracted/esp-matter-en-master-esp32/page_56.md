# Page 56

## Text Content

```
ESP-Matter Programming Guide, Release latest

5.3.2 Device Identity
Matter specification requires a unique Device Attestation Key (DAC) per device. This is a private ECDSA (secp256r1
curve) key that establishes the device identity to the Matter Ecosystem. DAC private needs to be protected from remote
as well as physical attacks in the best possible way.
Recommended ways for DAC private key protection:
• DAC private key can be protected using 5.2.2 Flash Encryption or 5.3.1 Secure Storage schemes.
Important: Support for DAC private key protection mechanisms described above is available in the Matter crypto port
layer for ESP32 platform.

Note: Espressif provides pre-provisioning service to build Matter-Compatible devices. This service also ensures the
security of the DAC private key and configuration data. Please contact Espressif Sales for more information.

1.5.4 5.4 More Security Considerations
Please refer to the overall ESP-IDF Security Guide for more considerations related to the debug interfaces, network,
transport and OTA updates related security.

1.5.5 5.5 Security Policy
The ESP-Matter GitHub repository has attached Security Policy Brief.
5.5.1 Advisories
• Espressif publishes critical Security Advisories, which includes security advisories regarding both hardware and
software.
• The specific advisories of the ESP-Matter software components shall be published through the GitHub repository.
5.5.2 Software Updates
Critical security issues in the ESP-Matter components, ESP-IDF components and dependant third-party libraries are
fixed as and when we find them or when they are reported to us. Gradually, we make the fixes available in all applicable
release branches in ESP-Matter.
Important: We recommend periodically updating to the latest bugfix version of the ESP-Matter release to have all
critical security fixes available.

52

Chapter 1. Table of Contents


```

