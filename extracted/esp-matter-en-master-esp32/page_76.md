# Page 76

## Text Content

```
ESP-Matter Programming Guide, Release latest

9.1.19 Keypad Input Cluster

Delegate Class

Reference Implementation

Keypad Input

Keypad Input Delegate

Delegate Class

Reference Implementation

Mode Select

Mode Select Delegate

9.1.20 Mode Select Cluster

9.1.21 Water Heater Management Cluster

Delegate Class

Reference Implementation

Water Heater Management

Water Heater Management Delegate

9.1.22 Energy Preference Cluster

Delegate Class

Reference Implementation

Energy Preference

Energy Preference Delegate

9.1.23 Commissioner Control Cluster

Delegate Class

Reference Implementation

Commissioner Control

Commissioner Control Delegate

Note:
Make sure that after implementing delegate class, you set the delegate class pointer at the time of creating
cluster.
robotic_vacuum_cleaner::config_t rvc_config;
rvc_config.rvc_run_mode.delegate = object_of_delegate_class;
endpoint_t *endpoint = robotic_vacuum_cleaner::create(node, & rvc_config, ENDPOINT_
,→FLAG_NONE);

72

Chapter 1. Table of Contents


```

