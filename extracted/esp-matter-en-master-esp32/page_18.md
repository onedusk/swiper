# Page 18

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.2.1.2 Configuring the Environment
This should be done each time a new terminal is opened
cd esp-idf; source ./export.sh; cd ..
cd esp-matter; source ./export.sh; cd ..

Enable Ccache for faster IDF builds.
Ccache is a compiler cache. Matter builds are very slow and takes a lot of time. Ccache caches the previous compilations
and speeds up recompilation in subsequent builds.
export IDF_CCACHE_ENABLE=1

Above can also be added to your shell’s profile file (.profile, .bashrc, .zprofile, etc.) to enable ccache every time you open
a new terminal.
2.2.2 ESP Matter Component (experimental)
You can check the component in Espressif Component Registry.
To add the esp_matter component to your project, run:
idf.py add-dependency "espressif/esp_matter^1.4.0"

An example with esp_matter component is offered:
• Managed Component Light
Note: To use this component, the version of IDF component management should be 1.4.* or >= 2.0. Use compote
version to show the version. Use pip install 'idf-component-manager~=1.4.0' or pip install
'idf-component-manager~=2.0.0' to install.

2.2.3 Building Applications
• Light
• Light Switch
• Zap Light
• Zigbee Bridge
• BLE Mesh Bridge

14

Chapter 1. Table of Contents


```

