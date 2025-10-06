# Page 36

## Text Content

```
ESP-Matter Programming Guide, Release latest

• Read attributes commands
• Read events commands
• Write attributes commands
• Subscribe attributes commands
• Subscribe events commands
• Group settings command.
2.10.1 Device console
Once you have flashed the controller example onto the device, you can use the device console to commission the device and
send commands to the end-device. All of the controller commands begin with the prefix matter esp controller.
2.10.2 Pairing commands
The pairing commands are used for commissioning end-devices and are available when the Enable matter
commissioner option is enabled. Here are three standard pairing methods:
• Onnetwork pairing: Prior to executing this commissioning method, it is necessary to connect both the controller
and the end-device to the same network and ensure that the commissioning window of the end-device is open.
To complete this process, you can use the command matter esp wifi connect. After the devices are
connected, the pairing process can be initiated.
matter esp wifi connect <ssid> <password>
matter esp controller pairing onnetwork <node_id> <setup_passcode>

• Ble-wifi pairing: This pairing method is supported for ESP32S3. Before you execute this commissioning method,
connect the controller to the Wi-Fi network and ensure that the end-device is in commissioning mode. You can
use the command matter esp wifi connect to connect the controller to your wifi network. Then we can
start the pairing.
matter esp wifi connect <ssid> <password>
matter esp controller pairing ble-wifi <node_id> <ssid> <password>
,→<pincode> <discriminator>

• Ble-thread pairing: This pairing method is supported for ESP32S3. Before you execute this commissioning
method, connect the controller to the Wi-Fi network in which there is a Thread Border Router (BR). And please
ensure that the end-device is in commissioning mode. You can use the command matter esp wifi connect
to connect the controller to your Wi-Fi network. Get the dataset tlvs of the Thread network that the Thread BR is
in. Then we can start the pairing.
matter esp wifi connect <ssid> <password>
matter esp controller pairing ble-thread <node_id> <dataset_tlvs> <pincode>
,→<discriminator>

• Matter payload based pairing: This method is similar to the previously mentioned pairing methods, but instead
of accepting a PIN code and discriminator, it uses a Matter setup payload as input. The setup payload is parsed to
extract the necessary information, which then initiates the pairing process.
For the code pairing method, commissioner tries to discover the end-device only on the IP network. However, when using
code-wifi, code-thread, or code-wifi-thread, and id CONFIG_ENABLE_ESP32_BLE_CONTROLLER
is enabled, controller tries to discover the end-device on both the IP and BLE networks.
Below are supported commands:
32

Chapter 1. Table of Contents


```

