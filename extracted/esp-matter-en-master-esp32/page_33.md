# Page 33

## Text Content

```
ESP-Matter Programming Guide, Release latest

cd esp-matter
export PATH=$PATH:$PWD/connectedhomeip/connectedhomeip/out/host

Generate the factory partition, please use the APPROPRIATE values for -v (Vendor Id), -p (Product Id), and -cd
(Certification Declaration).
esp-matter-mfg-tool --passcode 89674523 \
--discriminator 2245 \
-cd TEST_CD_FFF1_8001.der \
-v 0xFFF1 --vendor-name Espressif \
-p 0x8001 --product-name Bulb \
--hw-ver 1 --hw-ver-str DevKit

Few important output lines are mentioned below. Please take a note of onboarding codes, these can be used for commissioning the device.
[2022-12-02 11:18:12,059] [
[2022-12-02 11:18:12,059] [

INFO] - Generated QR code: MT:-24J06PF150QJ850Y10
INFO] - Generated manual code: 20489154736

Factory partition binary will be generated at the below path. Please check for <uuid>.bin file in this directory.
[2022-12-02 11:18:12,381] [
INFO] - Generated output files at: out/fff1_8001/
,→e17c95e1-521e-4979-b90b-04da648e21bb

2.7.4 Flashing firmware, secure cert and factory partition
Flashing secure cert partition. Please check partition table for esp_secure_cert partition address.
Note: Flash only if not flashed on manufacturing line.
esptool.py -p (PORT) write_flash 0xd000 secure_cert_partition.bin

Flashing factory partition, Please check the CONFIG_CHIP_FACTORY_NAMESPACE_PARTITION_LABEL for factory partition label. Then check the partition table for address and flash at that address.
esptool.py -p (PORT) write_flash 0x10000 path/to/partition/generated/using/mfg_tool/
,→uuid.bin

Flash application
idf.py flash

2.7.5 Test commissioning using chip-tool
If using the DACs signed by custom PAA that is not present in connectedhomeip repository, then download the PAA
certificate, please make sure it is in DER format.
Run the following command from host to commission the device.
./chip-tool pairing ble-wifi 1234 my_SSID my_PASSPHRASE my_PASSCODE my_DISCRIMINATOR ,→-paa-trust-store-path /path/to/PAA-Certificates/

1.2. 2. Developing with the SDK

29


```

