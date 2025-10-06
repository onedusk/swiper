# Page 41

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.10.8 Attestation Trust Storage
The controller example offers two options for the Attestation Trust Storage which is used to store and utilize the PAA
certificates for the Device Attestation verification. This feature is available when the Enable matter commissioner option is enabled in menuconfig. You can modify this setting in menuconfig Components -> ESP Matter
Controller -> Attestation Trust Store
• Attestation Trust Store - Test
Use two hardcoded PAA certificates(Chip-Test-PAA-FFF1-Cert&Chip-Test-PAA-NoVID-Cert) in the firmware.
• Attestation Trust Store - Spiffs
Read the PAA root certificates from the spiffs partition. The PAA der files should be placed in paa_cert
directory so that they can be flashed into the spiffs partition of the controller.

1.2.11 2.11 Custom Cluster
Matter enables users to implement custom clusters for unique features. This section introduces how to add a custom
cluster.
2.11.1 Cluster XML Template
Before adding a custom cluster, you should design the attributes, commands, and events it will include, and create the
cluster XML template file based on your design.
Example:
<?xml version="1.0"?>
<configurator>
<domain name="CHIP"/>
<cluster>
<domain>General</domain>
<name>Sample ESP</name>
<!-- The MSB 16 bits of <code> are the VendorID. Replace this with your
VendorID. 0x131B is the VendorId of Espressif.
The LSB 16 bits of <code> are a self-assigned ClusterID -->
<code>0x131BFC20</code>
<define>SAMPLE_ESP_CLUSTER</define>
<description>The Sample ESP cluster showcases a manufacturer custom cluster</
,→description>
<!-- Attributes -->
<!-- A simple test boolean attribute -->
<attribute side="server" code="0x0000" define="SAMPLE_BOOLEAN" type="boolean"␣
,→writable="true" default="false" optional="false">SampleBoolean</attribute>
<attribute side="server" code="0x0001" define="SAMPLE_CHAR_STR" type="char_string
,→" writable="false" optional="false">SampleCharStr</attribute>
<!-- Commands -->
<command source="client" code="0x00" name="CommandwithoutArgs" optional="false">
<description>
Simple command without any parameters and without a response.
</description>
</command>
(continues on next page)

1.2. 2. Developing with the SDK

37


```

