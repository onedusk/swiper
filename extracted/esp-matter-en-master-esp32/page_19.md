# Page 19

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.2.4 Flashing the Firmware
Choose IDF target.
• If IDF target has not been set explicitly, then esp32 is considered as default.
• The default device for esp32/esp32c3 is esp32-devkit-c/esp32c3-devkit-m. If you want to use another device, you can export ESP_MATTER_DEVICE_PATH after choosing the correct target, e.g. for m5stack
device: export ESP_MATTER_DEVICE_PATH=/path/to/esp_matter/device_hal/device/
m5stack
– If the device that you have is of a different revision, and is not working as expected, you can create a new
device and export your device path.
– The other peripheral components like led_driver, button_driver, etc. are selected based on the device selected.
– The configuration of the peripheral components can be found in $ESP_MATTER_DEVICE_PATH/
esp_matter_device.cmake.
(When flashing the SDK for the first time, it is recommended to do idf.py erase_flash to wipe out entire flash
and start out fresh.)
idf.py flash monitor

Note: If you are getting build errors like:
ERROR: This script was called from a virtual environment, can not create a virtual␣
,→environment again

It can be fixed by running below command:
pip install -r $IDF_PATH/requirements.txt

1.2.3 2.3 Commissioning and Control
There are a few implementations of Matter commissioners present in the connectedhomeip repository.
CHIP Tool is an example implementation of Matter commissioner and used for development purposes. An in-depth guide
on how to use chip-tool can be found in the CHIP Tool User Guide.
Espressif’s ESP RainMaker iOS and Android applications support commissioning and control of Matter devices.
• ESP-RainMaker Android App
• ESP-RainMaker iOS App

1.2. 2. Developing with the SDK

15


```

