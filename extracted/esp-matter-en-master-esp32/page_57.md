# Page 57

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.6 6. Configuration options to optimize RAM and Flash
1.6.1 6.1 Overview
There are several configuration options available to optimize Flash and RAM storage. The following list highlights key
options that significantly increase the free DRAM, heap, and reduce the flash footprint.
For more optimizations, we’ve also listed the reference links to esp-idf’s optimization guide.

1.6.2 6.2 Configurations
6.2.1 Test Environment setup
All numbers mentioned below are collected in the following environment:
Note:
• These numbers may vary slightly in a different environment.
• All numbers are in bytes
• As we are using BLE only for commissioning, BLE memory is freed post commissioning, hence there is an increase
in the free heap post commissioning. (CONFIG_USE_BLE_ONLY_FOR_COMMISSIONING=y)
• After building an example, some DRAM will be utilized, and the remaining DRAM will be allocated as heap.
Therefore, a direct increase in the free DRAM will reflect as an increase in free heap.

6.2.2 Default Configuration
We have used the default light example here, and below listed are the static and dynamic sizes.
6.2.3 Disable the chip-shell
Console shell is helpful when developing/debugging the application, but may not be necessary in production. Disabling
the shell can save space. Disable the below configuration option.
CONFIG_ENABLE_CHIP_SHELL=n

6.2.4 Adjust the dynamic endpoint count
The default dynamic endpoint count and default device type count is 16, which may be excessive for a normal application
creating only 2 endpoints. eg: light, only has two endpoints, one for root endpoint and one for actual light. Adjusting this
to a lower value, corresponding to the actual number of endpoints the application will create, can save DRAM.
Here, we have set the dynamic endpoint count and device type count to 2. Increase in the DRAM per endpoint/count is
~550 bytes.
CONFIG_ESP_MATTER_MAX_DYNAMIC_ENDPOINT_COUNT=2
CONFIG_ESP_MATTER_MAX_DEVICE_TYPE_COUNT=2

1.6. 6. Configuration options to optimize RAM and Flash

53


```

