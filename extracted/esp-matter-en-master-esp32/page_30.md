# Page 30

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.2.6 2.6 Factory Data Providers
2.6.1 Providers Introduction
There are four factory data providers, each with its own implementation, that need to be configured. These providers
supply the device with necessary factory data, which is then read by the device according to their respective implementations.
• Commissionable Data Provider
This particular provider is responsible for retrieving commissionable data, which includes information such as
setup-discriminator, spake2p-iteration-count, spake2p-salt, spake2p-verifier, and setup-passcode.
• Device Attestation Credentials(DAC) Provider
This particular provider is responsible for retrieving device attestation credentials, which includes information such
as CD, firmware-information, DAC, and PAI certificate. And it can also sign message with the DAC private key.
• Device Instance Info Provider
This particular provider is responsible for retrieving device instance information, which includes vendor-name,
vendor-id, product-name, product-id, product-url, product-label, hardware-version-string, hardware-version,
rotating-device-id-unique-id, serial-number, manufacturing-data, and part-number.
• Device Info Provider
This particular provider is responsible for retrieving device information, which includes fixed-labels, user-labels,
supported-locales, and supported-calendar-types.
2.6.2 Configuration Options
Different implementations of the four providers can be chosen in meuconfig:
• Commissionable Data Provider options in → Component config → ESP Matter
– Commissionable Data - Test, the device will use the hardcoded Commissionable Data. This
uses the legacy commissionable data provider and provides the test values. These test values are enclosed in
CONFIG_ENABLE_TEST_SETUP_PARAMS option and enabled by default.
– Commissionable
Data
Factory, the device will use commissionable
data information from the factory partition.This option is visible only when CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER is selected.
– Commissionable
Data
Secure
Cert, the device will use commissionable
data information from the secure cert partition.
This option is only visible when CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER and CONFIG_SEC_CERT_DAC_PROVIDER
is enabled.
– Commissionable
Data
Custom, the device will use the custom defined
commissionable
data
provider
to
obtain
commissionable
data
information.
esp_matter::set_custom_commissionable_data_provider() should be called before esp_matter::start() to set the custom provider.
Note: If you are using Factory, Secure Cert or Custom commissionable data provides, then disable the CONFIG_ENABLE_TEST_SETUP_PARAMS option.
• DAC Provider options in → Component config → ESP Matter
– Attestation - Test, the device will use the hardcoded Device Attestation Credentials.
26

Chapter 1. Table of Contents


```

