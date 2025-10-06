# Page 48

## Text Content

```
ESP-Matter Programming Guide, Release latest

3.3.2 Using ota_image_tool script
We should also edit the PROJECT_VER and the PROJECT_VER_NUMBER in the project’s CMakelists when using the
script to generate the OTA image.
• Build the example and generate the OTA image
cd path/to/example
idf.py build
cd path/to/esp_matter/connectedhomeip/connectedhomeip/src/app
./ota_image_tool.py create -v <vendor-id> -p <product-id> -vn 2 -vs v1.1 -da sha256 \
/path/to/original_app_bin /path/to/out_ota_bin

Note: The -vn (version-number) and -vs (version-string) should match the values in the project’s CMakelists.

1.3.4 3.4 PICS files
The PICS files define the Matter features for the product. The authorized test provider will determine the test cases to be
tested in Matter Certification Test according to the PICS files submitted.
The PICS Tool website is the tool to open, modify, validate, and save the XML PICS files. The reference XML PICS
template files include all the reference PICS files and each of the XML files defines the features of one or several clusters
on the products.
A PICS-generator tool is provided to generate the PICS files with the reference PICS XML template files. The tools will
read the supported clusters, attributes, commands, and event from a paired device and generate PICS files for that device.
Note that the Base XML file will not be generated with this tool. You still need to modify it in the PICS TOOL.
Open the reference PICS files that include all the clusters of the product, and select the features supported by the product.
Clicking the button Validate All, the PICS Tool will validate all the XML files and generate a list of test cases to
be tested in Matter Certification Test.

1.3.5 3.5 Route Information Option (RIO) notes
For Wi-Fi products using LwIP, TC-SC-4.9 should be tested in order to verify that the product can receive Router
Advertisement (RA) message with RIO and add route table that indicates whether the prefix can be reached by way of
the router. It can be tested with a Thread Border Router (BR) which sends RA message periodically and a Thread End
Device that is used to verify the Wi-Fi product can reach the Thread network via Thread BR. Some Wi-Fi Routers might
have the issue that they cannot forward RA message sent by the Thread BR, so please use a Wi-Fi Router that can forward
RA message when you are testing TC-SC-4.9.
Here are the steps to set up the Thread BR and Thread End Device. You should prepare 2 Radio Co-Processors (RCP)
to set up the ot-br-posix and ot-cli-posix. The RCP on ESP32-H2 is suggested to be used here. And you can also use
other platforms (such as nrf52840, efr32, etc.) as the RCPs.

44

Chapter 1. Table of Contents


```

