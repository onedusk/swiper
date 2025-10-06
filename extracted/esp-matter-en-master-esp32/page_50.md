# Page 50

## Text Content

```
ESP-Matter Programming Guide, Release latest

> dataset set active <PROVIDE THE DATASET OF THE BR THAT YOU NEED TO JOIN>
> dataset commit active
> ifconfig up
> thread start
> srp client autostart enable

In the console of ot-cli, discover the product IP address.
> dns service 177AC531F48BE736-0000000000000190 _matter._tcp.default.service.arpa.
DNS service resolution response for 177AC531F48BE736-0000000000000190 for service _
,→matter._tcp.default.service.arpa.
Port:5540, Priority:0, Weight:0, TTL:6913
Host:72FF282E7739731F.default.service.arpa.
HostAddress:fd11:66:0:0:22ae:27fe:13ac:54df TTL:6915
TXT:[SII=35303030, SAI=333030, T=30] TTL:6913

Note:
177AC531F48BE736-0000000000000190 can be get with command avahi-browse -rt
_matter._tcp. 177AC531F48BE736 is the compressed Fabric ID and 0000000000000190 is the node ID.
Ping the IP address of the Wi-Fi device.
> ping fd11:66:0:0:22ae:27fe:13ac:54df
16 bytes from fd11:66:0:0:22ae:27fe:13ac:54df : icmp_seq=2 hlim=64 time=14ms
1 packets transmitted, 1 packets received. Packet loss = 0.0%. Round-trip min/avg/max␣
,→= 14/14.0/14 ms.
Done

The ping command should be successful.

1.3.6 3.6 FW/SDK configuration notes
• Enable OTA Requestor in → Component config → CHIP Core → System Options
The option to enable OTA requestor. This option should be enabled if the OTA requestor feature is selected in
PICS files.
• Enable Extended discovery Support in → Component config → CHIP Device Layer
→ General Options
This option should be enabled if the PICS option MCORE.DD.EXTENDED_DISCOVERY is selected.
• Enable Device type in commissionable node discovery in → Component config →
CHIP Device Layer → General Options
This option should be enabled if the PICS option MCORE.SC.EXTENDED_DISCOVERY is selected.
• LOG_DEFAULT_LEVEL in → Component config → Log output
It is suggested to set log level to No output for passing the test cases of OnOff, LevelControl, and ColorControl
clusters. Here is related issue.

46

Chapter 1. Table of Contents


```

