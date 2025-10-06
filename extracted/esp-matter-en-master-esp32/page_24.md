# Page 24

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.5.1.3 Device Drivers
• The drivers, depending on the device, are typically initialized and updated in the example’s app_driver.cpp.
esp_err_t app_driver_init()
{
ESP_LOGI(TAG, "Initialising driver");
/* Initialize button */
button_config_t button_config = button_driver_get_config();
button_handle_t handle = iot_button_create(&button_config);
iot_button_register_cb(handle, BUTTON_PRESS_DOWN, app_driver_button_toggle_
,→cb);
app_reset_button_register(handle);
/* Initialize led */
led_driver_config_t led_config = led_driver_get_config();
led_driver_init(&led_config);
app_driver_attribute_set_defaults();
return ESP_OK;
}

• The driver’s attribute update API just handles the attributes that are actually relevant for the device. For example,
a color_temperature_light handles the power, brightness, hue, saturation and temperature.
esp_err_t app_driver_attribute_update(uint16_t endpoint_id, uint32_t cluster_id,␣
,→uint32_t attribute_id,
esp_matter_attr_val_t *val)
{
esp_err_t err = ESP_OK;
if (endpoint_id == light_endpoint_id) {
if (cluster_id == OnOff::Id) {
if (attribute_id == OnOff::Attributes::OnOff::Id) {
err = app_driver_light_set_power(val);
}
} else if (cluster_id == LevelControl::Id) {
if (attribute_id == LevelControl::Attributes::CurrentLevel::Id) {
err = app_driver_light_set_brightness(val);
}
} else if (cluster_id == ColorControl::Id) {
if (attribute_id == ColorControl::Attributes::CurrentHue::Id) {
err = app_driver_light_set_hue(val);
} else if (attribute_id ==␣
,→ColorControl::Attributes::CurrentSaturation::Id) {
err = app_driver_light_set_saturation(val);
} else if (attribute_id ==␣
,→ColorControl::Attributes::ColorTemperature::Id) {
err = app_driver_light_set_temperature(val);
}
}
}
return err;
}

20

Chapter 1. Table of Contents


```

