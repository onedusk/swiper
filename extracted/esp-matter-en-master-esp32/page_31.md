# Page 31

## Text Content

```
ESP-Matter Programming Guide, Release latest

– Attestation - Factory, the device will use the Device Attestation Credentials in the factory partition
binary. This option is visible only when CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER is
selected.
– Attestation - Secure Cert, the device will use the Device Attestation Credentials in the secure
cert partition. This option is for the Pre-Provisioned Modules. And the original vendor ID and product ID
should be added to the CD file for the Pre-Provisioned Modules. Please contact your Espressif contact person
for more information.
– Attestation - Custom, the device will use the custom defined DAC provider to obtain the Device Attestation Credentials. esp_matter::set_custom_dac_provider() should be called before esp_matter::start() to set the custom provider.
• Device Instance Info Provider options in → Component config → ESP Matter
– Device Instance Info - Test, the device will use the hardcoded Device Instance Information.
– Device Instance Info - Factory, the device will use device instance information from the factory
partition. This option is visable only when CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER
and ENABLE_ESP32_DEVICE_INSTANCE_INFO_PROVIDER is selected.
– Device Instance Info - Secure Cert, the device will use the unique identifier for generating
the rotating device identifier from the secure cert partition and all other details will be read from the factory
partition. This option is only visible when CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER
and CONFIG_SEC_CERT_DAC_PROVIDER is enabled.
– Device
Instance
Info
Custom, the device will use custom defined Device Instance Info Provider to obtain the Device Instance Information.
esp_matter::set_custom_device_instance_info_provider should be called before
esp_matter::start() to set the custom provider.
• Device Info Provider options in → Component config → ESP Matter
– Device Info - None, the device will not use any device information provider. It should be selected
when there are not related clusters on the device.
– Device Info - Factory, the device will use device information from the factory partition.
This option is visable only when CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER and ENABLE_ESP32_DEVICE_INFO_PROVIDER is selected.
– Device Info - Custom, the device will use custom defined Device Info Provider to obtain the Device Information. esp_matter::set_custom_device_info_provider should be called before
esp_matter::start() to set the custom provider.
2.6.3 Custom Providers
In order to use custom providers, you need to define implementations of the four base classes of the providers and override
the functions within them. And the custom providers should be set before esp_matter::start() is called.

1.2. 2. Developing with the SDK

27


```

