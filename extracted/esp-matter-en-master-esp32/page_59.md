# Page 59

## Text Content

```
ESP-Matter Programming Guide, Release latest

CONFIG_ESP_SYSTEM_EVENT_QUEUE_SIZE=16
CONFIG_ESP_SYSTEM_EVENT_TASK_STACK_SIZE=2048
CONFIG_MAX_EVENT_QUEUE_SIZE=20

Reduce the chip device event queue size can reduce IRAM size usage, lead to free heap increase.
CONFIG_MAX_EVENT_QUEUE_SIZE=20

6.2.8 Relocate certain code from IRAM to flash memory
Relocating certain code from IRAM to flash can reduce IRAM usage, so increase available heap size. However, this may
increase execution time.
Note: The options in this section may impact performance. Please perform thorough testing before using them in
production.

6.2.8.1 Reduce BLE IRAM usage
Move most IRAM into flash. This will increase the usage of flash and reduce ble performance. Because the code is moved
to the flash, the execution speed of the code is reduced. To have a small impact on performance, you need to enable flash
suspend (SPI_FLASH_AUTO_SUSPEND).
CONFIG_BT_CTRL_RUN_IN_FLASH_ONLY=y

6.2.8.2 Place FreeRTOS functions into Flash
When enabled the selected Non-ISR FreeRTOS functions will be placed into Flash memory instead of IRAM. This saves
up to 8KB of IRAM depending on which functions are used.
CONFIG_FREERTOS_PLACE_FUNCTIONS_INTO_FLASH=y

6.2.8.3 Place non-ISR ringbuf functions into flash
Place non-ISR ringbuf functions (like xRingbufferCreate/xRingbufferSend) into flash. This frees up IRAM, but the
functions can no longer be called when the cache is disabled.
CONFIG_RINGBUF_PLACE_FUNCTIONS_INTO_FLASH=y

1.6. 6. Configuration options to optimize RAM and Flash

55


```

