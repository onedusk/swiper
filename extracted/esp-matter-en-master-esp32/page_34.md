# Page 34

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.2.8 2.8 Matter OTA
• Enable the CONFIG_ENABLE_OTA_REQUESTOR option to enable Matter OTA Requestor functionality.
Please follow the OTA guide in the connectedhomeip repository for generating a Matter OTA image and performing
OTA.
2.8.1 Encrypted Matter OTA
The esp-matter SDK supports using a pre-encrypted application image for OTA upgrades. Please follow the steps below
to enable and use encrypted application images for OTA upgrades.
• Enable the CONFIG_ENABLE_OTA_REQUESTOR and CONFIG_ENABLE_ENCRYPTED_OTA options
• The application code must make an API call to esp_matter_ota_requestor_encrypted_init() after calling esp_matter::start(). You can use the following code as a reference:
#include <esp_matter_ota.h>
{
const char *rsa_private_key;
3072 private key in PEM format
uint16_t rsa_private_key_len;
,→private key

// Please set this to the buffer containing RSA␣

,→

,→

// Please set this to the length of RSA 3072␣

esp_err_t err = esp_matter_ota_requestor_encrypted_init(rsa_private_key, rsa_
private_key_len);

}

• Please refer to the encrypted OTA guide in the connectedhomeip repository for instructions on how to generate a
private key, encrypted OTA image, and Matter OTA image.
Note: There are several ways to store the private key, such as hardcoding it in the firmware, embedding it as a text file,
or reading it from the NVS. We have demonstrated the use of the private key by embedding it as a text file in the light
example.

1.2.9 2.9 Mode Select
This cluster provides an interface for controlling a characteristic of a device that can be set to one of several predefined
values. For example, the light pattern of a disco ball, the mode of a massage chair, or the wash cycle of a laundry machine.
2.9.1 Attribute Supported Modes
This attribute is the list of supported modes that may be selected for the CurrentMode attribute. Each item in this list
represents a unique mode as indicated by the Mode field of the ModeOptionStruct. Each entry in this list SHALL have
a unique value for the Mode field. ESP_MATTER uses factory partition to set the values of Supported Modes attribute.

30

Chapter 1. Table of Contents


```

