# Page 47

## Text Content

```
ESP-Matter Programming Guide, Release latest

• Automated Tests (Verified by UI-Automated and UI-SemiAutomated)
Here are the steps to upload the PAA certificate and use it for automated tests:
In Test Harness, you should modify the project configuration to use the vendor’s PAA for the DUT that requires a PAA
certificate to perform a pairing operation. The flag chip_tool_use_paa_certs in the dut_config should be
set to true to configure the Test Harness to use the PAA certificates.
"dut_config": {
"discriminator": "3840",
"setup_code": "20202021",
"pairing_mode": "onnetwork",
"chip_tool_timeout": null,
"chip_tool_use_paa_certs": true
}

Make sure to copy your PAA certificates in DER format to the default path /var/paa-root-certs/ on the
Raspberry-Pi.
sudo cp /path/to/PAA_certificate.der /var/paa-root-certs/

Run automated chip-tool tests and verify that the pairing commands are using the --paa-trust-store-path
option.
3.2.3 Menuconfig Options
Please consult the factory data providers and adjust the menucofig options accordingly for the certification test.

1.3.3 3.3 Matter OTA Image Generation
If the product supports OTA Requestor features of Matter, the test cases of OTA Software Update should be tested. So
you need to provide the image for OTA test and also the way to downgrade.
Here are two ways to generate the OTA image.
3.3.1 Using menuconfig option
Enable Generate Matter OTA image in → Component config → CHIP Device Layer → Matter OTA Image, set Device Vendor Id and Device Product Id in → Component config →
CHIP Device Layer → Device Identification Options, and edit the PROJECT_VER and the
PROJECT_VER_NUMBER in the project’s CMakelists. Build the example and the OTA image will be generated in the
build path with the app binary file.
Note: The PROJECT_VER_NUMBER must always be incremental. It must be higher than the version number of
firmware to be updated.

1.3. 3. Matter Certification

43


```

