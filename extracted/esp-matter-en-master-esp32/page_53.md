# Page 53

## Text Content

```
ESP-Matter Programming Guide, Release latest

• Spake2+ parameters: work as a proof of possession.
These details are generally programmed in the manufacturing partition that is unique per device. ESP-Matter provides a
utility (esp-matter-mfg-tool) to create these partition images on a per-device basis for mass manufacturing purposes.
When using the utility, by default, the above details will be included in the generated manufacturing partition image. The
utility also has a provision to include additional details in the same image by using CSV files.
Details about using the mass manufacturing utility can be found here: esp-matter-mfg-tool
4.3.2 Pre-Provisioned Modules
ESP32 modules can be pre-flashed with the manufacturing partition images during module manufacturing itself and then
be shipped to you.
This saves you the overhead of securely generating, encrypting and then programming the partition into the device at
your end.
Please contact your Espressif contact person for more information.
4.3.3 The esp-matter-mfg-tool Example
In Espressif Matter Prep-provisioning modules, the DAC key pair, DAC and PAI certificates are pre-flashed by default.
This section gives some examples on how to generate factory partition binary which includes :
Device unique data (Discriminator, Verifier, Serial Number, etc)
Manufacturing information (Vendor name, Product name, Hardware version, etc)
Note: The items listed in the examples are all mandatory, some common manufacturing information could be removed
if they are hard coded in the firmware.
This is the example to generate factory images after pre-provisioning:
• Generate generic factory image
esp-matter-mfg-tool -cd ~/test_cert/CD/Chip-CD-131B-1000.der -v 0x131B -,→vendor-name ESP -p 0x1000 --product-name light --hw-ver 1 --hw-ver-str␣
,→v1.0 --mfg-date 2022-10-25 --passcode 19861989 --discriminator 601 -,→serial-num esp32c_dev3

• Generate multiple generic factory images
esp-matter-mfg-tool -n 10 -cd ~/test_cert/CD/Chip-CD-131B-1000.der -v␣
,→0x131B --vendor-name ESP -p 0x1000 --product-name light --hw-ver 1 --hw,→ver-str v1.0 --mfg-date 2022-10-25

• Generate factory image with rotating device unique identify
esp-matter-mfg-tool -cd ~/test_cert/CD/Chip-CD-131B-1000.der -v 0x131B -,→vendor-name ESP -p 0x1000 --product-name light --hw-ver 1 --hw-ver-str␣
,→v1.0 --mfg-date 2022-10-25 --passcode 19861989 --discriminator 601 -,→serial-num esp32c_dev3 --enable-rotating-device-id --rd-id-uid␣
,→c0398f4980b07c9460f71c5421e1a3c5

• Generate multiple factory images with csv and mcsv
1.4. 4. Production Considerations

49


```

