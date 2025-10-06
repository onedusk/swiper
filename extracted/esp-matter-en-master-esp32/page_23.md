# Page 23

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.5.1 Building a Color Temperature Lightbulb
A device is represented in Matter in terms of its data model. As a first step of building your product, you will have to
define the data model for your device. Matter has a standard set of device types already defined that you can use. Please
refer to the Espressif Matter Blog for clarity on the terms like endpoints, clusters, etc. that are used in this section.
2.5.1.1 Data Model
• Typically, the data model is defined in the example’s app_main.cpp. First off we start by creating a Matter node,
which is the root of the Data Model.
node::config_t node_config;
node_t *node = node::create(&node_config, app_attribute_update_cb, NULL);

• We will use the color_temperature_light standard device type in this case. All standard device types
are available in esp_matter_endpoint.h header file. Each device type has a set of default configuration that can be
specific as well.
color_temperature_light::config_t light_config;
light_config.on_off.on_off = DEFAULT_POWER;
light_config.level_control.current_level = DEFAULT_BRIGHTNESS;
endpoint_t *endpoint = color_temperature_light::create(node, &light_config,␣
,→ENDPOINT_FLAG_NONE);

In this case, we create the light using the color_temperature_light::create() function. Similarly,
multiple endpoints can be created on the same node. Check the following sections for more info.
2.5.1.2 Attribute Callback
• Whenever a Matter client makes changes to the device, they end up updating the attributes in the data model.
• When an attribute is updated, the attribute_update_cb is used to notify the application of this change. You
would typically call device driver specific APIs for executing the required action. Here, if the callback type is
PRE_UPDATE, the driver is updated first. If that is a success, only then the attribute value is actually updated in
the database.
esp_err_t app_attribute_update_cb(callback_type_t type, uint16_t endpoint_id,␣
,→uint32_t cluster_id,
uint32_t attribute_id, esp_matter_attr_val_t␣
,→*val, void *priv_data)
{
esp_err_t err = ESP_OK;
if (type == PRE_UPDATE) {
/* Driver update */
err = app_driver_attribute_update(endpoint_id, cluster_id, attribute_id,␣
,→val);
}
return err;
}

1.2. 2. Developing with the SDK

19


```

