# Page 71

## Text Content

```
ESP-Matter Programming Guide, Release latest

Classes

class custom_encodable_type : public EncodableToTLV
class multiple_write_encodable_type
class custom_command_callback : public chip::app::CommandSender::Callback

1.8 Enabling ESP-Insights in ESP-Matter
• To learn more about esp-insights and get started, please refer project README.md.
• Before building the app, enable the option ESP_INSIGHTS_ENABLED through menuconfig.
• Follow the steps present set up esp-insights account , and create an auth key.
• Create a file named insights_auth_key.txt in the project directory of the example.
• Download the auth key and copy Auth Key to the example.
cp /path/to/auth/key.txt path/to/esp-matter/examples/generic_switch/
,→insights_auth_key.txt

• Refer the esp-matter Generic Switch example to enable the traces and metrics reported by the esp32 tracing
backend in the chip SDK on the insights dashboard and about how to use the auth key for enabling insights.
• Enable the option ENABLE_ESP_INSIGHTS_SYSTEM_STATS to get a report of the system metrics in the chip SDK
on the insights dashboard.

1.9 9. Application User Guide
1.9.1 9.1. Delegate Implementation
As per the implementation in the connectedhomeip repository, some of the clusters require an application defined delegate
to consume specific data and actions. In order to provide this flexibity to the application, esp-matter facilitates delegate
initilization callbacks in the cluster create API. It is expected that application will define it’s data and actions in the form
of delegate-impl class and set the delegate while creating cluster/device type.
List of clusters with delegate:
• Mode Base Cluster (all derived types of clusters).
• Content Launch Cluster.
• Fan Control Cluster.
• Audio Output Cluster.
• Energy EVSE Cluster.
• Device Energy Management Cluster.
• Microwave Oven Control Cluster.
• Door Lock Cluster.

1.8. Enabling ESP-Insights in ESP-Matter

67


```

