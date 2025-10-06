# Page 37

## Text Content

```
ESP-Matter Programming Guide, Release latest

matter esp controller pairing code <node_id> <setup_payload>
matter esp controller pairing code-wifi <node_id> <ssid> <passphrase> <setup_
,→payload>
matter esp controller pairing code-thread <node_id> <operationalDataset>
,→<setup_payload>
matter esp controller pairing code-wifi-thread <node_id> <ssid> <passphrase>
,→<operationalDataset> <setup_payload>

2.10.3 Cluster commands
The invoke-cmd command is used for sending cluster commands to the end-devices. It utilizes a cluster_command class to establish the sessions and send the command packets. The class constructor function could
accept two callback inputs:
• Success callback: This callback will be called upon the reception of the success response. It could be used to handle
the response data for the command that requires a reponse. Now the default success callback will print the response
data for GroupKeyManagement, Groups, Scenes, Thermostat, and DoorLock clusters. If you want to handle the
response data in your example, you can register your success callback when creating the cluster_command
object.
• Error callback: This callback will be called upon the reception of the failure response or reponse timeout.

• Send the cluster command:
matter esp controller invoke-cmd <node-id | group-id> <endpoint-id>
,→<cluster-id> <command-id> <command-data>

Note:
• The command-data should utilize a JSON object string and the name of each item in this object should be \
"<TagNumber>:<DataType>\" or \"<TagName>:<TagNumber>:<DataType>\". The TagNumber should be the same as the command parameter ID in Matter SPEC and the supported DataTypes are listed in
$ESP_MATTER_PATH/components/esp_matter/utils/json_to_tlv.h
• For the DataType bytes, the value should be a Base64-Encoded string.
Here are some examples of the command-data format.
• For MoveToLevel command in LevelControl cluster, the command-data ({"level": 10, "transitionTime": 0, "optionsMask": 0, "optionsOverride": 0}) should be:
matter esp controller invoke-cmd <node-id> <endpoint-id> 8 0 "{\"0:U8\": 10, \
,→"1:U16\": 0, \"2:U8\": 0, \"3:U8\": 0}"

• For
KeySetWrite
command
in
GroupKeyManagement
cluster,
the
command-data
({"groupKeySet":{"groupKeySetID":
42,
"groupKeySecurityPolicy":
0,
"epochKey0": d0d1d2d3d4d5d6d7d8d9dadbdcdddedf, "epochStartTime0": 2220000,
"epochKey1":
null, "epochStartTime1":
null, "epochKey2":
null,
"epochStartTime2": null}}) should be:

1.2. 2. Developing with the SDK

33


```

