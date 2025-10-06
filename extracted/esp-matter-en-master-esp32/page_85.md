# Page 85

## Text Content

```
ESP-Matter Programming Guide, Release latest

1.11.17 A1.17 Updating attribute marked as ATTRIBUTE_FLAG_MANAGED_INTERNALLY
When an attribute is marked with the flag ATTRIBUTE_FLAG_MANAGED_INTERNALLY, application can not directly
modify the attribute’s value using esp_matter::attribute::update. To update such attributes, retrieve the
corresponding delegate or instance within the cluster implementation and perform the update through it. For example, to
update DefaultOTAProviders attribute in OTASoftwareUpdateRequestor cluster, use the following code:
chip::OTARequestorInterface * request = chip::GetRequestorInstance();
if (request) {
chip::OTARequestorInterface::ProviderLocationType provider;
provider.providerNodeID = 123;
provider.endpoint = 0;
provider.fabricIndex = 1;
request->AddDefaultOtaProvider(provider);
}

1.11. A1 Appendix FAQs

81


```

