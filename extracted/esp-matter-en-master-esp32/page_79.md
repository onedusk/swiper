# Page 79

## Text Content

```
ESP-Matter Programming Guide, Release latest

– If you have gdb enabled, run the command backtrace and share the output of gdb too.

1.11.4 A1.4 Device not crashed but not responding
My device is not responding to commands:
• Make sure your device is commissioned successfully and is connected to the Wi-Fi.
• Make sure the node_id and the endpoint_id are correct in the command from chip-tool.
• If you are still facing issues, reproduce it on the default example for the device and then raise a Github issue. Along
with the details mentioned in the issue template, please share the following details:
– The steps you followed to reproduce the issue.
– The complete device logs taken over UART.

1.11.5 A1.5 Onboard LED not working
The LED on my devkit is not working:
• Make sure you have selected the proper device.
ESP_MATTER_DEVICE_PATH to the correct path.

You can explicitly do that by exporting the

• Check the version of your board, and if it has the LED connected to a different pin. If it is different, you can change
the led_driver_config_t accordingly in the device.c file.
• If you are still facing issues, reproduce it on the default example for the device and then raise a Github issue.

1.11.6 A1.6 Using Rotating Device Identifier
What is Rotating Device Identifier:
• The Rotating Device Identifier provides a non-trackable identifier which is unique per-device and that can be used
in one or more of the following ways:
– Provided to the vendor’s customer support for help in pairing or establishing Node provenance;
– Used programmatically to obtain a Node’s Passcode or other information in order to provide a simplified
setup flow. Note that the mechanism by which the Passcode may be obtained is outside of this specification.
If the Rotating Device Identifier is to be used for this purpose, the system implementing this feature SHALL
require proof of possession by the user at least once before providing the Passcode. The mechanism for this
proof of possession, and validation of it, is outside of this specification.
How to use Rotating Device Identifier
• Enable the Rotating Device Identifier support in menuconfig.
• Add the --enable-rotating-device-id and add the --rd-id-uid to specify the Rotating ID
Unique ID when use the esp-matter-mfg-tool to generate partition.bin file.
Difference between Rotating ID Unique ID and Unique ID
• The Rotating ID Unique ID is a parameter used to generate Rotating Device Identifier, it is
a unique per-device identifier and shall consist of a randomly-generated 128-bit or longer octet string which shall
be programmed during factory provisioning or delivered to the device by the vendor using secure means after a
software update, it shall stay fixed during the lifetime of the device.

1.11. A1 Appendix FAQs

75


```

