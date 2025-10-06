# Page 17

## Text Content

```
ESP-Matter Programming Guide, Release latest

cd esp-idf
source ./export.sh
cd ..
git clone --depth 1 https://github.com/espressif/esp-matter.git
cd esp-matter
git submodule update --init --depth 1
cd ./connectedhomeip/connectedhomeip
./scripts/checkout_submodules.py --platform esp32 linux --shallow
cd ../..
./install.sh
cd ..

• For Mac OS-X host:
cd esp-idf
source ./export.sh
cd ..
git clone --depth 1 https://github.com/espressif/esp-matter.git
cd esp-matter
git submodule update --init --depth 1
cd ./connectedhomeip/connectedhomeip
./scripts/checkout_submodules.py --platform esp32 darwin --shallow
cd ../..
./install.sh
cd ..

Note: The modules for platform linux or darwin are required for the host tools building.

Note:
If you don’t want to install host tools (chip-tool, chip-cert etc.)
--no-host-tool.

you can use ./install.sh

To clone the esp-matter repository with all the submodules, use the following command:
cd esp-idf
source ./export.sh
cd ..
git clone --recursive https://github.com/espressif/esp-matter.git
cd esp-matter
./install.sh
cd ..

Note: If it runs into some errors like:
dial tcp 108.160.167.174:443: connect: connection refused
ConnectionResetError: [Errno 104] Connection reset by peer

It’s probably caused by some network connectivity issue, a VPN is required for most of the cases.

1.2. 2. Developing with the SDK

13


```

