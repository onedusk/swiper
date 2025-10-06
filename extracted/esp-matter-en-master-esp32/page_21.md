# Page 21

## Text Content

```
ESP-Matter Programming Guide, Release latest

https://project-chip.github.io/connectedhomeip/qrcode.html?data=MT:Y.K9042C00KA0648G00

If you want to use different values for commissioning the device, please use the esp-matter-mfg-tool to generate the
factory partition which has to be flashed on the device. It also generates the new pairing code and QR code image using
which you can commission the device.
2.3.1.2 Post Commissioning Setup
The device would need additional configuration depending on the example, for it to work. Check the “Post Commissioning
Setup” section in examples for more information.
• Light
• Light Switch
• Zap Light
• Zigbee Bridge
• BLE Mesh Bridge
2.3.1.3 Cluster Control
Use the cluster commands to control the attributes.
onoff toggle 0x7283 0x1
onoff on 0x7283 0x1
levelcontrol move-to-level 10 0 0 0 0x7283 0x1
levelcontrol move-to-level 100 0 0 0 0x7283 0x1
colorcontrol move-to-color-temperature 0 10 0 0 0x7283 0x1

chip-tool when used in interactive mode uses CASE resumption as against establishing CASE for cluster control commands. This results into shorter execution times, thereby improving the overall experience.
For more details about the commands, please check chip-tool usage guide

1.2.4 2.4 Device console
The console on the device can be used to run commands for testing. It is configurable through menuconfig and enabled
by default in the firmware. Here are some useful commands:
• BLE commands: Start and stop BLE advertisement:
matter ble [start|stop|state]

• Wi-Fi commands: Set and get the Wi-Fi mode:
matter wifi mode [disable|ap|sta]

• Device configuration: Dump the device static configuration:
1.2. 2. Developing with the SDK

17


```

