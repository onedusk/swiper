# Page 78

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.11 A1 Appendix FAQs
1.11.1 A1.1 Compilation errors
I cannot build the application:
• Make sure you are on the correct esp-idf branch/release.
• Run git submodule update —init —recursive to make sure the submodules are at the correct heads.
• Make sure you have the correct ESP_MATTER_PATH (and any other required paths).
• Delete the build/ directory and also sdkconfig and sdkconfig.old and then build again.
• If you are still facing issues, reproduce it on the default example and then raise a Github issue.

1.11.2 A1.2 Device commissioning using chip-tool
I cannot commission a new device through the chip-tool:
• If the chip-tool pairing ble-wifi command is failing, make sure the arguments are correct.
• Please check chip-tool pairing ble-wifi --help for argument help.
• Make sure Bluetooth is turned on, on your client (host).
Bluetooth/BLE does not work on by device:
• There is a known issues #13303 where BLE does not work on MacOS.
• In this case, the following can be done:
– Run the device console command: matter esp wifi connect <ssid> <password>.
– Run the chip-tool command for commissioning over ip: chip-tool pairing onnetwork 0x7283
20202021.
• If you are still facing issues, reproduce it on the default example for the device and then raise a Github issue.

1.11.3 A1.3 Device crashing
My device is crashing:
• Given the tight footprint requirements of the device, please make sure any issues in your code have been ruled out.
If you believe the issue is with the Espressif SDK itself, please recreate the issue on the default example application
(without any changes) and go through the following steps:
• Make sure you are on the correct esp-idf branch. Run git submodule update —init —recursive to
make sure the submodules are at the correct heads.
• Make sure you have the correct ESP_MATTER_PATH (and any other paths) is (are) exported.
• Delete the build/ directory and also sdkconfig and sdkconfig.old and then build and flash again.
• If you are still facing issues, reproduce it on the default example for the device and then raise a Github issue. Along
with the details mentioned in the issue template, please share the following details:
– The steps you followed to reproduce the issue.
– The complete device logs taken over UART.
– The .elf file from the build/ directory.

74

Chapter 1. Table of Contents


```

