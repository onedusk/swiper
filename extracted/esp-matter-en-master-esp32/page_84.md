# Page 84

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.11.15 A1.15 Using BLE after Matter commissioning
Most Matter applications do not require BLE after commissioning. By default, BLE is deinitialized after commissioning
to reclaim RAM and increase the available free heap. Refer to A1.8 Why does free RAM increase after first commissioning
for more details.
However, if BLE functionality is needed even after commissioning, you can disable the CONFIG_USE_BLE_ONLY_FOR_COMMISSIONING option. This ensures that the memory allocated to BLE functionality
is retained, allowing BLE to be used for other purposes post-commissioning.
After commissioning is complete, Matter will stop advertising, but the application can utilize BLE for other roles or
operations. e.g. BLE Peripheral, BLE Central, etc.
To learn more, refer to the bleprph and blecent examples in esp-idf/examples/bluetooth/nimble. These
examples demonstrate BLE Peripheral and BLE Central roles. It also provides the step-by-step tutorial for building such
devices.
For implementation details on Peripheral and Central roles, refer to the bleprph_advertise() and blecent_scan() functions
in the respective examples.
BLE Central role is disabled by default in the esp-matter SDK’s default example configurations. Please enable CONFIG_BT_NIMBLE_ROLE_CENTRAL option if you plan to use BLE Central role.
Note: Above mentioned details apply specifically to the NimBLE host.
For more advanced BLE usage, you can use the external platform feature. It also serves as a way to integrate custom BLE
usage with Matter.
Please refer to the advance setup section in the programming guide. This has been demonstrated in the blemesh_bridge
examples.

1.11.16 A1.16 Moving BSS Segments to PSRAM to Reduce Memory Usage
The BSS section of libesp_matter.a and libCHIP.a can consume significant internal memory. For devices with PSRAM,
you can move the BSS segments to external memory to significantly reduce the internal memory footprint.
To move the BSS segments of libCHIP.a and libesp_matter.a into external RAM:
1. Enable the CONFIG_ESP_ALLOW_BSS_SEG_EXTERNAL_MEMORY option in menuconfig.
2. Create a linker.lf file in your project’s main component, you can check the the example linker.lf file.
3. Modify your main component’s CMakeLists.txt to include:
set(ldfragments linker.lf)
idf_component_register(
...
LDFRAGMENTS "${ldfragments}")

This configuration will move the BSS segments to PSRAM when CONFIG_ESP_ALLOW_BSS_SEG_EXTERNAL_MEMORY
is enabled, significantly reducing the internal memory usage of your application.
Please check #1123 for relevant discussion on Github issue.

80

Chapter 1. Table of Contents


```

