# Page 25

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.5.2 Defining your own data model
This section demonstrates creating standard endpoints, clusters, attributes, and commands that are defined in the Matter
specification
2.5.2.1 Endpoints
The device can be customized by editing the endpoint/device_type creating in the app_main.cpp of the example. Examples:
• on_off_light:
on_off_light::config_t light_config;
endpoint_t *endpoint = on_off_light::create(node, &light_config, ENDPOINT_FLAG_
,→NONE);

• fan:
fan::config_t fan_config;
endpoint_t *endpoint = fan::create(node, &fan_config, ENDPOINT_FLAG_NONE);

• door_lock:
door_lock::config_t door_lock_config;
endpoint_t *endpoint = door_lock::create(node, &door_lock_config, ENDPOINT_FLAG_
,→NONE);

• window_covering_device:
window_covering_device::config_t window_covering_device_config(static_cast<uint8_
,→t>
,→(chip::app::Clusters::WindowCovering::EndProductType::kTiltOnlyInteriorBlind));
endpoint_t *endpoint = window_covering_device::create(node, &window_covering_
,→config, ENDPOINT_FLAG_NONE);

The window_covering_device config_t structure includes a constructor that allows specifying an end
product type different than the default one, which is “Roller shade”. Once a config_t instance has been instantiated, its end product type cannot be modified.
• pump
pump::config_t pump_config(1, 10, 20);
endpoint_t *endpoint = pump::create(node, &pump_config, ENDPOINT_FLAG_
,→NONE);

The pump config_t structure includes a constructor that allows specifying maximum pressure, maximum speed and maximum flow values. If they aren’t set, they will be set to null by default. Once a
config_t instance has been instantiated, these three values cannot be modified.

1.2. 2. Developing with the SDK

21


```

