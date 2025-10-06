# Page 26

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.5.2.2 Clusters
Additional clusters can also be added to an endpoint. Examples:
• on_off:
on_off::config_t on_off_config;
cluster_t *cluster = on_off::create(endpoint, &on_off_config, CLUSTER_FLAG_
,→SERVER);

• temperature_measurement:
temperature_measurement::config_t temperature_measurement_config;
cluster_t *cluster = temperature_measurement::create(endpoint, &temperature_
,→measurement_config, CLUSTER_FLAG_SERVER);

• window_covering:
window_covering::config_t window_covering_config(static_cast<uint8_
,→t>
,→(chip::app::Clusters::WindowCovering::EndProductType::kTiltOnlyInteriorBlind));
,→

window_covering_config.feature_flags = window_
,→covering::feature::lift::get_id();
cluster_t *cluster = window_covering::create(endpoint, &window_
,→covering_config, CLUSTER_FLAG_SERVER);

The window_covering config_t structure includes a constructor that allows specifying an end
product type different than the default one, which is “Roller shade”. Once a config_t instance has
been instantiated, its end product type cannot be modified.
• pump_configuration_and_control:
pump_configuration_and_control::config_t pump_configuration_and_control_
,→config(1, 10, 20);
pump_configuration_and_control_config..feature_flags = pump_configuration_
,→and_control::feature::constant_pressure::get_id();
cluster_t *cluster = pump_configuration_and_control::create(endpoint, &
,→pump_configuration_and_control_config, CLUSTER_FLAG_SERVER);

The pump_configuration_and_control config_t structure includes a constructor that allows specifying maximum pressure, maximum speed and maximum flow values. If they aren’t set, they
will be set to null by default. Once a config_t instance has been instantiated, these three values
cannot be modified.
2.5.2.3 Attributes and Commands
Additional attributes and commands can also be added to a cluster. Examples:
• attribute: on_off:
bool default_on_off = true;
attribute_t *attribute = on_off::attribute::create_on_off(cluster, default_on_
,→off);

• attribute: cluster_revision:

22

Chapter 1. Table of Contents


```

