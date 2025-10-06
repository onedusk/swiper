# Page 35

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.9.2 Generate Factory Partition Using esp-matter-mfg-tool
Use esp-matter-mfg-tool to generate factory partition of the supported modes attribute.
2.9.2.1 Usage
esp-matter-mfg-tool -cn "My bulb" -v 0xFFF2 -p 0x8001 --pai \
-k path/to/esp-matter/connectedhomeip/connectedhomeip/credentials/test/attestation/
,→Chip-Test-PAI-FFF2-8001-Key.pem \
-c path/to/esp-matter/connectedhomeip/connectedhomeip/credentials/test/attestation/
,→Chip-Test-PAI-FFF2-8001-Cert.pem \
-cd path/to/esp-matter/connectedhomeip/connectedhomeip/credentials/test/certification,→declaration/Chip-Test-CD-FFF2-8001.der \
--supported-modes mode1/label1/endpointId/"value\mfgCode, value\mfgCode" mode2/
,→label2/endpointId/"value\mfgCode, value\mfgCode"

• For empty Semantic Tags list
--supported-modes mode1/label1/endpointId

mode2/label2/endpointId

2.9.3 Build example
For example we want to use mode_select cluster in light example.
• Add source and include path to example/light/main/CMakeList.txt
Append "${MATTER_SDK_PATH}/examples/platform/esp32/mode-support" to SRC_DIRS and PRIV_
,→INCLUDE_DIRS

• In file example/light/app_main.cpp.
#include <static-supported-modes-manager.h>
ModeSelect::StaticSupportedModesManager sStaticSupportedModesManager;
{
cluster::mode_select::config_t ms_config;
cluster_t *ms_cluster = cluster::mode_select::create(endpoint, &ms_config,␣
,→CLUSTER_FLAG_SERVER, ESP_MATTER_NONE_FEATURE_ID);
sStaticSupportedModesManager.InitEndpointArray(get_count(node));
ModeSelect::setSupportedModesManager(&sStaticSupportedModesManager);
}

1.2.10 2.10 Matter Controller
This section introduces the Matter controller example. Now this example supports the following features of the standard
Matter controller:
• BLE-WiFi pairing
• BLE-Thread pairing
• On-network pairing
• Invoke cluster commands
1.2. 2. Developing with the SDK

31


```

