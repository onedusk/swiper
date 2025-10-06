# Page 80

## Text Content

```
ESP-Matter Programming Guide, Release latest

• The Unique ID is an attribute in Basic Information Cluster, it shall indicate a unique identifier for
the device, which is constructed in a manufacturer specific manner. It may be constructed using a permanent device
identifier (such as device MAC address) as basis. In order to prevent tracking:
– it SHOULD NOT be identical to (or easily derived from) such permanent device identifier
– it SHOULD be updated when the device is factory reset
– it SHALL not be identical to the SerialNumber attribute
– it SHALL not be printed on the product or delivered with the product

1.11.7 A1.7 ModuleNotFoundError: No module named ‘lark’
Encountering the above error while building the esp-matter example could indicate that the steps outlined in the getting
the repository section of the documentation were not properly followed.
The esp-matter example relies on several python dependencies that can be found in the requirements.txt . These dependencies must be installed into the python environment of the esp-idf to ensure that the example builds successfully.
One recommended approach to installing these requirements is by running the command source $IDF_PATH/
export.sh before running esp-matter/install.sh, as suggested in the programming guide. However, if the
error persists, you can try the following steps to resolve it:
cd esp-idf
source ./export.sh
cd esp-matter
python3 -m pip install -r requirements.txt
# Now examples will build without any error
cd examples/...
idf.py build

1.11.8 A1.8 Why does free RAM increase after first commissioning
After the first commissioning, you may notice that the free RAM increases. This is because, by default, BLE is only used
for the commissioning process. Once the commissioning is complete, BLE is deinitialized, and all the memory allocated
to it is recovered. Here’s the link to the implementation which frees the BLE memory .
However, if you want to continue using the BLE even after the commissioning process, you can disable the CONFIG_USE_BLE_ONLY_FOR_COMMISSIONING. This will ensure that the memory allocated to the BLE functionality
is not released after the commissioning process, and the free RAM won’t go up.

1.11.9 A1.9 How to generate Matter Onboarding Codes (QR Code and Manual Pairing
Code)
When creating a factory partition using esp-matter-mfg-tool, both the QR code and manual pairing codes are
generated.
Along with that, there are two more methods for generating Matter onboarding codes:
• Python script: generate_setup_payload.py

76

Chapter 1. Table of Contents


```

