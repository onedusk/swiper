# Page 27

## Text Content

```
ESP-Matter Programming Guide, Release latest

uint16_t default_cluster_revision = 1;
attribute_t *attribute = global::attribute::create_cluster_revision(cluster,␣
,→default_cluster_revision);

• command: toggle:
command_t *command = on_off::command::create_toggle(cluster);

• command: move_to_level:
command_t *command = level_control::command::create_move_to_level(cluster);

2.5.2.4 Features
Mandatory features for a device type or endpoint can be configured at endpoint level.
• feature: lighting: On/Off cluster:
extended_color_light::config_t light_config;
light_config.on_off_lighting.start_up_on_off = nullptr;
endpoint_t *endpoint = extended_color_light::create(node, &light_config, ENDPOINT_
,→FLAG_NONE, nullptr);

Few of some mandatory feature for a cluster (i.e. cluster having O.a/O.a+ feature conformance) can be configured at
cluster level. For example: Thermostat cluster has O.a+ conformance for Heating and Cooling features, that means at
least one of them should be present on the thermostat cluster while creating it.
• feature: heating: Thermostat cluster:
thermostat::config_t thermostat_config;
thermostat_config.features.heating.occupied_heating_setpoint = 2200;
thermostat_config.feature_flags = thermostat::feature::heating::get_id();
cluster::thermostat::create(endpoint, &(config->thermostat_config), CLUSTER_FLAG_
,→SERVER);

Optional features which are applicable to a cluster can also be added.
• feature: taglist: Descriptor cluster:
cluster_t* cluster = cluster::get(endpoint, Descriptor::Id);
descriptor::feature::taglist::add(cluster);

2.5.3 Adding custom data model fields
This section demonstrates creating custom endpoints, clusters, attributes, and commands that are not defined in the Matter
specification and can be specific to the vendor.

1.2. 2. Developing with the SDK

23


```

