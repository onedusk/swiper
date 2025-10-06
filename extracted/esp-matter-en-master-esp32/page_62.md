# Page 62

## Text Content

```
ESP-Matter Programming Guide, Release latest

(continued from previous page)

CONFIG_SUPPORT_TIME_SYNCHRONIZATION_CLUSTER=n
CONFIG_SUPPORT_TIMER_CLUSTER=n
CONFIG_SUPPORT_TVOC_CONCENTRATION_MEASUREMENT_CLUSTER=n
CONFIG_SUPPORT_UNIT_TESTING_CLUSTER=n
CONFIG_SUPPORT_USER_LABEL_CLUSTER=n
CONFIG_SUPPORT_VALVE_CONFIGURATION_AND_CONTROL_CLUSTER=n
CONFIG_SUPPORT_WAKE_ON_LAN_CLUSTER=n
CONFIG_SUPPORT_LAUNDRY_WASHER_CONTROLS_CLUSTER=n
CONFIG_SUPPORT_LAUNDRY_DRYER_CONTROLS_CLUSTER=n
CONFIG_SUPPORT_WIFI_NETWORK_MANAGEMENT_CLUSTER=n
CONFIG_SUPPORT_WINDOW_COVERING_CLUSTER=n
CONFIG_SUPPORT_WATER_HEATER_MANAGEMENT_CLUSTER=n
CONFIG_SUPPORT_WATER_HEATER_MODE_CLUSTER=n

Table 1: Static memory stats

Used D/IRAM
Used Flash

Size

Decreased by

179487
1576436

3736
36938

Table 2: Dynamic memory stats

On Bootup
Post Commissioning

Free Heap

Increased by

44256
77976

3876
4164

1.6.3 6.3 References for futher optimizations
• RAM optimization
• Binary size optimization
• Speed Optimization
• ESP32 Memory Analysis — Case Study
• Optimizing IRAM can provide additional heap area but at the cost of execution speed. Relocating frequently-called
functions from IRAM to flash may result in increased execution time

1.7 7. API Reference
1.7.1 Data Model
This has the high level APIs for Data Model.

58

Chapter 1. Table of Contents


```

