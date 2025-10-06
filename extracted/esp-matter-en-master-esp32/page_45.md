# Page 45

## Text Content

```
ESP-Matter Programming Guide, Release latest

3.2.1 Certification Declaration
A Certification Declaration (CD) is a cryptographic document that allows a Matter device to assert its protocol compliance.
It can be generated with following steps. We need to generate the CD which matches the vendor id and product id in
DAC and the ones in basic information cluster.
A test CD signed by the test CD signing keys in connectedhomeip SDK repository is required for Matter Certification Test,
so the certification_type of it is 1 (provisional). The CD in official products passing the Matter Certification
Test is issued by CSA and the certification_type is 2 (official).
• Generate the Test CD file
cd path/to/esp_matter/connnectedhomeip/connnectedhomeip
out/host/chip-cert gen-cd --format-version 1 --vendor-id 0x131B --product-id 0x1234 \
--device-type-id 0x010c --certificate-id CSA00000SWC00000,→01 \
--security-level 0 --security-info 0 --version-number 1 \
--certification-type 1 \
--key credentials/test/certification-declaration/Chip-Test,→CD-Signing-Key.pem \
--cert credentials/test/certification-declaration/Chip-Test,→CD-Signing-Cert.pem \
--out path/to/test_CD_file

Note:
• The option --certification-type must be 1 for the Matter Certification Test.
• The options --vendor_id (vendor_id) should be the Vendor ID (VID) that the vendor receives from CSA, and
--product_id (product_id) could be the Product ID (PID) choosed by the vendor. They should be the same
as the attributes’ value in basic information cluster.
• If the product uses the DACs and PAI certifications provided by a trusted third-party certification authority, the VID and PID in DAC are different from the ones in basic information cluster. Then the
--dac-origin-vendor-id and --dac-origin-product-id options should be added in the command generating the test CD file.

3.2.2 Certificates and Keys
For Matter Certification Test, vendors should generate their own test Product Attestation Authority (PAA) certificate,
Product Attestation Intermediate (PAI) certificate, and Device Attestation Certificate (DAC), but not use the default test
PAA certificate in connectedhomeip SDK repository. So you need to generate a PAA certificate, and use it to sign and
attest PAI certificates which will be used to sign and attest the DACs. The PAI certificate, DAC, and DAC’s private key
should be stored in the product you submit to test.
Here are the steps to generate the certificates and keys using chip-cert and esp-matter-mfg-tool.

1.3. 3. Matter Certification

41


```

