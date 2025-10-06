# Page 44

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.11.2.4 Custom Cluster Functions
A custom cluster might requires special funtions to handle initialization, attribute changes, shutdown, and pre-attribute
changes. For instance, the AAI and CHI need to be registered so that they can be accessed by the Matter data model.
Therefore, the cluster requires an initialization function to register them. To enable these functions, the cluster should be
added to the appropriate entry in the zap configuration data file.
2.11.3 Example Usage
• Zap Example
If the example uses zap tool to generate its data model, the custom cluster should be added to the example’s zap
file. The zap tool will then generate the data model code, including the custom cluster, during the building process.
• ESP-Matter Example
If the example uses ESP-Matter APIs to define its data model, the custom data model should be created and added
to the data model using the esp-matter APIs, following the instructions in Adding custom data model fields

1.3 3. Matter Certification
The Matter Certification denotes compliance to a Connectivity Standards Alliance (CSA) specification for the product
and allow the use of Certified Product logos and listing of the product on the Alliance website for verification.
You need to become a member of CSA and request a Vendor ID code from CSA Certification before you apply for a
Matter Certification. Then you need to choose an authorized test provider (must be validated for Matter testing) and
submit your product for testing. Here are some tips for the Matter Certification Test.

1.3.1 3.1 Introduction to Test Harness (TH)
Test Harness on RaspberryPi is used for Matter Certification Test. You can fetch the TH RaspberryPi image from here
and install the image to a micro SD card with the Raspberry Pi Imager.
Test cases can be verified with TH by 4 methods including UI-Automated, UI-SemiAutomated, UI-Manual, and Verification Steps Document. A website UI is used for the first three methods. You can follow the instructions in TH User
Guide to use the website UI. For the last method, you should use the chip-tool in path ~/apps of the TH and execute
the commands in the Verification Steps Document step by step.

1.3.2 3.2 Matter Factory Partition Binary
Matter factory partition binary files contains the commissionable information (discriminator, salt, iteration count, and
spake2+ verifier) and device attestation information (Certification Declaration (CD), Product Attestation Intermediate
(PAI) certificate, Device Attestation Certificate (DAC), and DAC private key), device instance information(vendor ID,
vendor name, product ID, product name, etc.), and device information (fixed label, supported locales, etc.). These informations are used to identify the product and ensure the security of commissioning.

40

Chapter 1. Table of Contents


```

