# Page 82

## Text Content

```
ESP-Matter Programming Guide, Release latest

• Locking the Matter thread
lock::chip_stack_lock(portMAX_DELAY);
... // eg: access Matter attribute, open/close commissioning window.
lock::chip_stack_unlock();

• Scheduling the work on Matter thread
static void WorkHandler(intptr_t context);
{
... // Do the stuff
}
chip::DeviceLayer::PlatformMgr().ScheduleWork(WorkHandler, <intptr_t>
,→(nullptr));

1.11.11 A1.11 Firmware Version Number
Similar to the ESP-IDF’s application versioning scheme, the ESP-Matter SDK provides two options for setting the
firmware version. It depends on CONFIG_APP_PROJECT_VER_FROM_CONFIG option and by default option is
disabled.
If the CONFIG_APP_PROJECT_VER_FROM_CONFIG option is disabled, you need to set the version and version
string by defining the CMake variables in the project’s CMakeLists.txt file. All the examples use this scheme and
have these variables set. Here’s an example:
set(PROJECT_VER "1.0")
set(PROJECT_VER_NUMBER 1)

On the other hand, if the CONFIG_APP_PROJECT_VER_FROM_CONFIG option is enabled, you need to set the
version using the following configuration options:
• Software Version
Set the CONFIG_DEVICE_SOFTWARE_VERSION_NUMBER option. (Component config -> CHIP Device
Layer -> Device Identification Options -> Device Software Version Number)
• Software Version String
Set the CONFIG_APP_PROJECT_VER option. (Application manager -> Get the project version from
Kconfig)
Note:
• Ensure you use the correct versioning scheme when building the OTA image.
• Verify that the software version number in the firmware matches the one specified in the Matter OTA header.
• The software version number of the OTA image must be numerically higher.
• If you need to perform a functional rollback, the version number in the OTA image must be higher than the current
version, even though the binary content may match the previous OTA image.

78

Chapter 1. Table of Contents


```

