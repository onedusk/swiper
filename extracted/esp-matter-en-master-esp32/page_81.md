# Page 81

## Text Content

```
ESP-Matter Programming Guide, Release latest

./generate_setup_payload.py --discriminator 3131 --passcode 20201111 \
--vendor-id 65521 --product-id 32768 \
--commissioning-flow 0 --discovery-cap,→bitmask 2

• chip-tool
// Generate the QR Code
chip-tool payload generate-qrcode --discriminator 3131 --setup-pin-code␣
,→20201111 \
--vendor-id 0xFFF1 --product-id 0x8004 \
--version 0 --commissioning-mode 0 -,→rendezvous 2
// Generates the short manual pairing code (11-digit).
chip-tool payload generate-manualcode --discriminator 3131 --setup-pin,→code 20201111 \
--version 0 --commissioning-mode 0
// To generate a long manual pairing code (21-digit) that includes both␣
,→the vendor ID and product ID,
// --commissioning-mode parameter must be set to either 1 or 2,␣
,→indicating a non-standard commissioning flow.
chip-tool payload generate-manualcode --discriminator 3131 --setup-pin,→code 20201111 \
--vendor-id 0xFFF1 --product-id␣
,→0x8004 \
--version 0 --commissioning-mode 1

To create a QR code image, copy the QR code text and paste it into CHIP QR Code.

1.11.10 A1.10 Chip stack locking error … Code is unsafe/racy
E (84728) chip[DL]: Chip stack locking error at 'src/system/
,→SystemLayerImplFreeRTOS.cpp:55'. Code is unsafe/racy
E (84728) chip[-]: chipDie chipDie chipDie
abort() was called at PC 0x40139b7f on core 0
0x40139b7f:␣
,→chip::Platform::Internal::AssertChipStackLockedByCurrentThread(char const*,
,→ int) at /home/jonathan/Desktop/Workspace/firmware/build/esp-idf/chip/../..
,→/../../esp-matter/connectedhomeip/connectedhomeip/config/esp32/third_party/
,→connectedhomeip/src/lib/support/CodeUtils.h:508
(inlined by) chipDie at /home/jonathan/Desktop/Workspace/firmware/build/esp,→idf/chip/../../../../esp-matter/connectedhomeip/connectedhomeip/config/
,→esp32/third_party/connectedhomeip/src/lib/support/CodeUtils.h:518
(inlined by)␣
,→chip::Platform::Internal::AssertChipStackLockedByCurrentThread(char const*,
,→ int) at /home/jonathan/Desktop/Workspace/firmware/build/esp-idf/chip/../..
,→/../../esp-matter/connectedhomeip/connectedhomeip/config/esp32/third_party/
,→connectedhomeip/src/platform/LockTracker.cpp:36

When interacting with Matter resources, it is necessary to perform the operations from within the Matter thread to
avoid assertion errors. This applies to tasks such as getting and setting attributes, invoking commands, and performing
operations using the server’s object, such as opening or closing the commissioning window.
To address this, there are two possible approaches:

1.11. A1 Appendix FAQs

77


```

