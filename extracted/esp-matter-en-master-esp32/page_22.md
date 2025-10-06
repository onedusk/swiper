# Page 22

## Text Content

```
ESP-Matter Programming Guide, Release latest

matter config

• Factory reset:
matter esp factoryreset

• On-boarding codes: Dump the on-boarding pairing code payloads:
matter onboardingcodes

Additional Matter specific commands:
• Get attribute: (The IDs are in hex):
matter esp attribute get <endpoint_id> <cluster_id> <attribute_id>

– Example: on_off::on_off:
matter esp attribute get 0x1 0x6 0x0

• Set attribute: (The IDs are in hex):
matter esp attribute set <endpoint_id> <cluster_id> <attribute_id> <attribute␣
,→value>

– Example: on_off::on_off:
matter esp attribute set 0x1 0x6 0x0 1

• Diagnostics:
matter esp diagnostics mem-dump

• Wi-Fi
matter esp wifi connect <ssid> <password>

• Bridge device:
matter esp bridge <command>

– Example: add (Parent endpoint should have aggregator device type):
matter esp bridge add <parent_endpoint_id> <device_type_id>

1.2.5 2.5 Developing your Product
Understanding the structure before actually modifying and customising the device is helpful.

18

Chapter 1. Table of Contents


```

