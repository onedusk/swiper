# Page 28

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.5.3.1 Endpoints
Non-Standard endpoint can be created, without any clusters.
• Endpoint create:
endpoint_t *endpoint = endpoint::create(node, ENDPOINT_FLAG_NONE);

2.5.3.2 Clusters
Non-Standard/Custom clusters can also be created:
• Cluster create:
uint32_t custom_cluster_id = 0x131bfc00;
cluster_t *cluster = cluster::create(endpoint, custom_cluster_id, CLUSTER_FLAG_
,→SERVER);

2.5.3.3 Attributes and Commands
Non-Standard/Custom attributes can also be created on any cluster:
• Attribute create:
uint32_t custom_attribute_id = 0x0;
uint16_t default_value = 100;
attribute_t *attribute = attribute::create(cluster, custom_attribute_id,␣
,→ATTRIBUTE_FLAG_NONE, esp_matter_uint16(default_value);

• Command create:
static esp_err_t command_callback(const ConcreteCommandPath &command_path,␣
,→TLVReader &tlv_data, void
*opaque_ptr)
{
ESP_LOGI(TAG, "Custom command callback");
return ESP_OK;
}
uint32_t custom_command_id = 0x0;
command_t *command = command::create(cluster, custom_command_id, COMMAND_FLAG_
,→ACCEPTED, command_callback);

2.5.4 Advanced Setup
This section explains adding external platforms for Matter. This step is optional for most devices. Espressif’s SDK
for Matter provides support for overriding the default platform layer, so the BLE and Wi-Fi implementations can be
customized. Here are the required steps for adding an external platform layer.

24

Chapter 1. Table of Contents


```

