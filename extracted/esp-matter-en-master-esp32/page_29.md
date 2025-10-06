# Page 29

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.5.4.1 Creating the external platform directory
Create a directory platform/${NEW_PLATFORM_NAME} in your codebase.
You can typically copy
${ESP_MATTER_PATH}/connectedhomeip/connectedhomeip/src/platform/ESP32 as a start.
Note that the new platform name should be something other than ESP32. In this article we’ll use ESP32_custom as
an example. The directory must be under platform folder to meet the Matter include path conventions.
2.5.4.2 Modifying the BUILD.gn target
There is an example BUILD.gn file for the ESP32_custom example platform. It simply compiles the ESP32 platform
in Matter without any modifications.
• The new platform directory must be added to the Matter include path. See the ESP32_custom_include
config in the above mentioned file.
• Multiple build configs must be exported to the build system. See the buildconfig_header section in the file
for the required definitions.
2.5.4.3 Editing Kconfigs
• Enable CONFIG_CHIP_ENABLE_EXTERNAL_PLATFORM.
• Set CONFIG_CHIP_EXTERNAL_PLATFORM_DIR to the relative path from ${ESP_MATTER_PATH}/
connectedhomeip/connectedhomeip/config/esp32 to the external platform directory. For instance, if your source tree is:
my_project
├── esp-matter
└── platform
└── ESP32_custom

Then CONFIG_CHIP_EXTERNAL_PLATFORM_DIR
ESP32_custom.

would

be

../../../../../platform/

• Disable CONFIG_BUILD_CHIP_TESTS.
• If your external platform does not support the connectedhomeip/connectedhomeip/src/lib/shell/ provided in the Matter shell library, then disable CONFIG_ENABLE_CHIP_SHELL.
2.5.4.4 Example Usage
As an example, you can build light example on ESP32_custom platform with following steps:
mkdir $ESP_MATTER_PATH/../platform
cp -r $ESP_MATTER_PATH/connectedhomeip/connectedhomeip/src/platform/ESP32 $ESP_MATTER_
,→PATH/../platform/ESP32_custom
cp $ESP_MATTER_PATH/examples/common/external_platform/BUILD.gn $ESP_MATTER_PATH/../
,→platform/ESP32_custom
cd $ESP_MATTER_PATH/examples/light
cp sdkconfig.defaults.ext_plat sdkconfig.defaults
idf.py build

1.2. 2. Developing with the SDK

25


```

