# Page 58

## Text Content

```
ESP-Matter Programming Guide, Release latest

6.2.5 Use the newlib nano formatting
This optimization saves approximately 25-50K of flash, depending on the target. In our case, it results in a flash reduction
of 47 KB.
Additionally, it lowers the high watermark of task stack for functions that call printf() or other string formatting functions.
Fore more details please take a look at esp-idf’s newlib nano formatting guide.
CONFIG_NEWLIB_NANO_FORMAT=y

6.2.6 BLE Optimizations
Since most devices will primarily operate as BLE peripherals and typically won’t need more than one connection (especially if it’s just a Matter app), we can optimize by reducing the maximum allowed connections, thereby saving DRAM.
Additionally, given the peripheral nature of these devices, we can disable the central and observer roles, for further optimization. In current implementation, BLE is disabled once commissioning succeeds, so these optimizations do not
contribute to free heap post-commissioning.
Below are the configuration options that can be set to achieve these optimizations.
CONFIG_NIMBLE_MAX_CONNECTIONS=1
CONFIG_BTDM_CTRL_BLE_MAX_CONN=1
CONFIG_BT_NIMBLE_MAX_CONNECTIONS=1
CONFIG_BT_NIMBLE_ROLE_CENTRAL=n
CONFIG_BT_NIMBLE_ROLE_OBSERVER=n
CONFIG_BT_NIMBLE_MAX_BONDS=2
CONFIG_BT_NIMBLE_MAX_CCCDS=2
CONFIG_BT_NIMBLE_SECURITY_ENABLE=n
CONFIG_BT_NIMBLE_50_FEATURE_SUPPORT=n
CONFIG_BT_NIMBLE_WHITELIST_SIZE=1
CONFIG_BT_NIMBLE_GATT_MAX_PROCS=1
CONFIG_BT_NIMBLE_MSYS_1_BLOCK_COUNT=10
CONFIG_BT_NIMBLE_MSYS_1_BLOCK_SIZE=100
CONFIG_BT_NIMBLE_MSYS_2_BLOCK_COUNT=4
CONFIG_BT_NIMBLE_MSYS_2_BLOCK_SIZE=320
CONFIG_BT_NIMBLE_ACL_BUF_COUNT=5
CONFIG_BT_NIMBLE_HCI_EVT_HI_BUF_COUNT=5
CONFIG_BT_NIMBLE_HCI_EVT_LO_BUF_COUNT=3
CONFIG_BT_NIMBLE_ENABLE_CONN_REATTEMPT=n

6.2.7 Configuring logging event buffer
Matter events serve as a historical record, stored in chronological order in the logging event buffer. By reducing the buffer
size we can potentially save the DRAM. However, it’s important to note that this reduction could lead to the omission of
events.
For instance, reducing the critical log buffer from 4K to 256 bytes could save 3K+ DRAM, but it comes with the trade-off
of potentially missing critical events.
CONFIG_EVENT_LOGGING_CRIT_BUFFER_SIZE=256
CONFIG_EVENT_LOGGING_INFO_BUFFER_SIZE=256
CONFIG_EVENT_LOGGING_DEBUG_BUFFER_SIZE=256
CONFIG_MAX_EVENT_QUEUE_SIZE=20

Reduce ESP system event queue size and event task stack size can increase free heap size.

54

Chapter 1. Table of Contents


```

