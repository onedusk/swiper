# Page 46

## Text Content

```
ESP-Matter Programming Guide, Release latest

3.2.2.1 Generating PAA Certificate
Vendor scoped PAA certificate is suggested for the Matter Certificate Test. It can be generated with the help of blow
mentioned steps.
Generate the vendor scoped PAA certificate and key, please make sure to change the --subject-vid (vendor_id)
option base on the one that is being used.
cd path/to/connnectedhomeip/out/host/
./chip-cert gen-att-cert --type a --subject-cn "Example PAA CN" --subject-vid 0x131B \
--valid-from "2021-06-28 14:23:43" --lifetime 4294967295 \
--out-key /path/to/PAA_key \
--out /path/to/PAA_certificate

3.2.2.2 Generating Factory Partition Binary Files
After getting the PAA certificate and key, the factory partition binary files with PAI certificate, DAC, and DAC keys can
be generated using esp-matter-mfg-tool.
• Install the requirements and export the dependent tools path if not done already
cd path/to/esp_matter
python3 -m pip install -r requirements.txt
export PATH=$PATH:$PWD/connectedhomeip/connectedhomeip/out/host

• Generate factory partition binary files
esp-matter-mfg-tool -n <count> -cn Espressif --paa -c /path/to/PAA_certificate -k /
,→path/to/PAA_key \
-cd /path/to/CD_file -v 0x131B --vendor-name Espressif -p 0x1234 \
--product-name Test-light --hw-ver 1 --hw-ver-str v1.0

Note: For more information about the arguments, you can use esp-matter-mfg-tool --help
The option -n (count) is the number of generated binaries. In the above command, esp-matter-mfg-tool will generate PAI
certificate and key and then use them to generate count different DACs and keys. It will use the generated certificates
and keys to generate count factory partition binaries with different DACs, discriminators, and setup pincodes. Flash
the factory binary to the device’s NVS partition. Then the device will send the vendor’s PAI certificate and DAC to the
commissioner during commissioning.
3.2.2.3 Using Vendor’s PAA in Test Harness(TH)
• Manual Tests (Verified by UI-Manual and Verification Steps Document)
The option --paa-trust-store-path should be added when using chip-tool to pair the device for manual tests.
Note:
• pincode and discriminator are in the /out/<vid>-<pid>/<UUID>/<uuid>-onb_codes.csv.
• PAA certificate should be
paa-certificate-path.

42

converted

to

DER

format

using

chip-cert

and

stored

in

Chapter 1. Table of Contents


```

