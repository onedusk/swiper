# Page 60

## Text Content

```
ESP-Matter Programming Guide, Release latest

6.2.8.4 Use esp_flash implementation in ROM
Enable this flag to use new SPI flash driver functions from ROM instead of ESP-IDF. After enable CONFIG_SPI_FLASH_ROM_IMPL, will increase free IRAM. But may miss out on some flash features and support for
new flash chips.
CONFIG_SPI_FLASH_ROM_IMPL=y
CONFIG_SPI_MASTER_ISR_IN_IRAM=n
CONFIG_SPI_SLAVE_ISR_IN_IRAM=n

6.2.8.5 Force the entire heap component to be placed in flash memory
Enable this flag to save up RAM space by placing the heap component in the flash memory Note that it is only safe to
enable this configuration if no functions from esp_heap_caps.h or esp_heap_trace.h are called from IRAM ISR which
runs when cache is disabled.
CONFIG_HEAP_PLACE_FUNCTION_INTO_FLASH=y

6.2.9 Reduce Task Stack Size
Reduce some task stack size can increase free heap size.
CONFIG_ESP_MAIN_TASK_STACK_SIZE=3072
CONFIG_ESP_TIMER_TASK_STACK_SIZE=2048
CONFIG_CHIP_TASK_STACK_SIZE=6144

6.2.10 Excluding Unused Matter Clusters
If the cluster implementation source files use a class derived from another class with virtual functions and instantiate a
global object of this class, the linker may keep all the related symbols that may be used for this class in the vtable. To
eliminate these symbols, you can deselect the unused Matter clusters under → Component config → ESP Matter
→ Select Supported Matter Clusters. Excluding unused clusters will help reduce flash and memory usage.
The default configuration disables all unused clusters.
CONFIG_SUPPORT_ACCOUNT_LOGIN_CLUSTER=n
CONFIG_SUPPORT_ACTIVATED_CARBON_FILTER_MONITORING_CLUSTER=n
CONFIG_SUPPORT_AIR_QUALITY_CLUSTER=n
CONFIG_SUPPORT_APPLICATION_BASIC_CLUSTER=n
CONFIG_SUPPORT_APPLICATION_LAUNCHER_CLUSTER=n
CONFIG_SUPPORT_AUDIO_OUTPUT_CLUSTER=n
CONFIG_SUPPORT_BOOLEAN_STATE_CONFIGURATION_CLUSTER=n
CONFIG_SUPPORT_BRIDGED_DEVICE_BASIC_INFORMATION_CLUSTER=n
CONFIG_SUPPORT_CARBON_DIOXIDE_CONCENTRATION_MEASUREMENT_CLUSTER=n
CONFIG_SUPPORT_CARBON_MONOXIDE_CONCENTRATION_MEASUREMENT_CLUSTER=n
CONFIG_SUPPORT_CHANNEL_CLUSTER=n
CONFIG_SUPPORT_CHIME_CLUSTER=n
CONFIG_SUPPORT_COMMISSIONER_CONTROL_CLUSTER=n
CONFIG_SUPPORT_CONTENT_LAUNCHER_CLUSTER=n
CONFIG_SUPPORT_CONTENT_CONTROL_CLUSTER=n
CONFIG_SUPPORT_CONTENT_APP_OBSERVER_CLUSTER=n
CONFIG_SUPPORT_DEVICE_ENERGY_MANAGEMENT_CLUSTER=n
CONFIG_SUPPORT_DEVICE_ENERGY_MANAGEMENT_MODE_CLUSTER=n
(continues on next page)

56

Chapter 1. Table of Contents


```

