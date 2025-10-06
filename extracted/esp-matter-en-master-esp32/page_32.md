# Page 32

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.2.7 2.7 Using esp_secure_cert partition
2.7.1 Configuration Options
Build the firmware with below configuration options
# Disable the DS Peripheral support
CONFIG_ESP_SECURE_CERT_DS_PERIPHERAL=n
# Use DAC Provider implementation which reads attestation data from secure cert␣
,→partition
CONFIG_SEC_CERT_DAC_PROVIDER=y
# Enable some options which reads CD and other basic info from the factory partition
CONFIG_ENABLE_ESP32_FACTORY_DATA_PROVIDER=y
CONFIG_ENABLE_ESP32_DEVICE_INSTANCE_INFO_PROVIDER=y
CONFIG_FACTORY_COMMISSIONABLE_DATA_PROVIDER=y
CONFIG_FACTORY_DEVICE_INSTANCE_INFO_PROVIDER=y

2.7.2 Certification Declaration
If you do not have an certification declaration file then you can generate the test CD with the help of below mentioned
steps. We need to generate the new CD because it SHALL match the VID, PID in DAC and the ones reported by basic
cluster.
• Build the host tools if not done already
cd connectedhomeip/connectedhomeip
gn gen out/host
ninja -C build

Generate the Test CD, please make sure to change the -V (vendor_id) and -p (product-id) options based on the ones that
are being used. For more info about the arguments, please check chip-cert’s gen-cd command in the connectedhomeip
repository.
out/host/chip-cert gen-cd -f 1 -V 0xFFF1 -p 0x8001 -d 0x0016 \
-c "CSA00000SWC00000-01" -l 0 -i 0 -n 1 -t 0 \
-K credentials/test/certification-declaration/Chip-Test-CD,→Signing-Key.pem \
-C credentials/test/certification-declaration/Chip-Test-CD,→Signing-Cert.pem \
-O TEST_CD_FFF1_8001.der

2.7.3 Factory Partition
Factory partition contains basic information like VID, PID, etc.
By default, the CD(Certification Declaration) is stored in the factory partition and we need to add the -cd option when
generating the factory partition.
Alternatively, if you’d like to embed the CD in the firmware, you can enable the CONFIG_ENABLE_SET_CERT_DECLARATION_API option and use the SetCertificationDeclaration()
API to set the CD. You can refer to the reference implementation in :project_file: light example.
Export the dependent tools path

28

Chapter 1. Table of Contents


```

