# Page 83

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.11.12 A1.12 Stuck at “Solving dependencies requirements …..”
When building an example, if it is stuck at “Solving dependencies requirements…” you can resolve this issue by clearing
the component manager cache.
# On Linux
rm -rf ~/.cache/Espressif/ComponentManager
# On macOS
rm -rf ~/Library/Caches/Espressif/ComponentManager

1.11.13 A1.13 ESP32-C2 log garbled, unable to perform Matter commissioning and
other abnormal issues
When encountering the above issues, the following possible causes may exist: 1. Incorrect baud rate settings. See
UART console baud rate 2. Incorrect XTAL crystal frequency settings. The default XTAL crystal frequency in the
SDK examples is 26 Mhz, if the ESP32-C2 board used for testing is 40 MHz, please change the configuration as CONFIG_XTAL_FREQ_40=y. See Main XTAL frequency You can check the XTAL frequency with this command.
$ esptool.py flash_id
esptool.py v4.7.0
Serial port /dev/ttyUSB0
Connecting....
Detecting chip type... ESP32-C2
Chip is ESP32-C2 (revision v1.0)
Features: WiFi, BLE
Crystal is 26MHz
MAC: 08:3a:8d:49:b3:90

1.11.14 A1.14 Generating Matter Onboarding Codes on the device itself
The Passcode serves as both proof of possession for the device and the shared secret needed to establish the initial secure
channel for onboarding.
For best practices in Passcode generation and storage on the device, refer to Section 5.1.7: Generation of the Passcode
in the Core Matter Specification.
Ideally, devices should only store the Spake2p verifier, not the Passcode itself. If the Passcode is stored on the device, it
must be physically separated from the Spake2p verifier’s location and must be accessible only through local interface and
must not be accessible to the unit handling the Spake2p verifier.
For devices capable of displaying the onboarding payload, the use of a dynamic Passcode is recommended.
The Light Switch example in the SDK demonstrates the use of a dynamic Passcode. It implements a custom Commissionable Data Provider that generates the dynamic Passcode, along with the corresponding Spake2p verifier and onboarding
payload, directly on the device.
Please check #1128 and #1126 for relevant discussion on Github issue

1.11. A1 Appendix FAQs

79


```

