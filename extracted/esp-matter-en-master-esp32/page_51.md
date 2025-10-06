# Page 51

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.3.7 3.7 Appendix FAQs
Here are some issues that you might meet in Matter Certification Test and quick solutions for them.
• TC-CNET-3.11
No response on step 7 is expected (Related issue).
All the NetworkCommissioning commands are fail-safe required. If the commands fail with a FAILSAFE_REQUIRED status code. You need to send arm-fail-safe command and then send the NetworkCommissioning commands.
– TC-RR-1.1
For more application endpoints with group cluster, need more nvs size to store group table, so if
the TC-RR-1.1 failed, can try to increase the nvs size.
(Related issue <https://github.com/projectchip/connectedhomeip/issues/32481>`__)
Please note that the minimum NVS size required is 48 KB (0xC000) when using a single endpoint with a group
cluster.

1.4 4. Production Considerations
1.4.1 4.1 Prerequisites
All Matter examples use certain test or evaluation values that enables you to quickly build and test Matter. As you get ready
to go to production, these must be replaced with the actual values. These values are typically a part of the manufacturing
partition in your device.
4.1.1 Vendor ID and Product ID
A Vendor Identifier (VID) is a 16-bit number that uniquely identifies a particular product manufacturer or a vendor. It
is allocated by the Connectivity Standards Alliance (CSA). Please reach out to CSA for this.
A Product Identifier (PID) is a 16-bit number that uniquely identifies a product of a vendor. It is assigned by the vendor
(you).
A VID-PID combination uniquely identifies a Matter product.
4.1.2 Certificates
A Device Attestation Certificate (DAC) proves the authenticity of the device manufacturer and the certification status
of the device’s hardware and software. Every Matter device must have a DAC and corresponding private key, unique to
it. The device should also have a Product Attestation Intermediate (PAI) certificate that was used to sign and attest the
DAC. The PAI certificate in turn is signed and attested by Product Attestation Authority (PAA). The PAA certificate is
an implicitly trusted self-signed root certificate.
Please reach out to your Espressif representative for the details about how to procure the DAC.

1.4. 4. Production Considerations

47


```

