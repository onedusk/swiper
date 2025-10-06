# Page 54

## Text Content

```
ESP-Matter Programming Guide, Release latest

esp-matter-mfg-tool -cd ~/test_cert/CD/Chip-CD-131B-1000.der -v 0x131B -,→vendor-name ESP -p 0x1000 --product-name light --hw-ver 1 --hw-ver-str␣
,→v1.0 --enable-rotating-device-id --mfg-date 2022-10-25 --csv mfg.csv -,→mcsv mfg_m.csv

• The example of csv and mcsv file
• CSV:
serial-num,data,string
rd-id-uid,data,hex2bin
discriminator,data,u32
• MCSV:
serial-num,rd-id-uid,discriminator
esp32c_dev3,c0398f4980b07c9460f71c5421e1a3c5,1234
esp32c_dev4,c0398f4980b07c9460f71c5421e1a3c6,1235
esp32c_dev5,c0398f4980b07c9460f71c5421e1a3c7,1236
esp32c_dev6,c0398f4980b07c9460f71c5421e1a3c8,1237
esp32c_dev7,c0398f4980b07c9460f71c5421e1a3c9,1238
4.3.4 Recommended Providers to Use

Note: WARNING: These options are not recommended for devices that are already in field or modules that reads data
from the factory partition or some other source.
We recommend using the following providers:
• Commissionable data provider: secure cert
• Device attestation data provider: secure cert
• Device instance info provider: secure cert
Below are the configuration options that should be enabled. These can be appended to sdkconfig.defaults.
In the following example, we demonstrate a different approach that places the configurations in a separate file, which is
then used with the idf.py build command.
cat > sdkconfig.defaults.prod <<EOF
# Enable the implementations in the connectedhomeip repo
CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER=y
CONFIG_ENABLE_ESP32_DEVICE_INSTANCE_INFO_PROVIDER=y
# Set the appropriate providers
CONFIG_SEC_CERT_DAC_PROVIDER=y
CONFIG_SEC_CERT_COMMISSIONABLE_DATA_PROVIDER=y
CONFIG_SEC_CERT_DEVICE_INSTANCE_INFO_PROVIDER=y
CONFIG_NONE_DEVICE_INFO_PROVIDER=y
EOF
idf.py -D SDKCONFIG_DEFAULTS="sdkconfig.defaults.prod" set-target esp32c3 build

50

Chapter 1. Table of Contents


```

