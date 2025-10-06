# Page 39

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.10.5 Write attribute commands
The write-attr command is used for sending the commands of writing attributes on the end-device.
• Send the write-attribute command:
matter esp controller write-attr <node-id> <endpoint-id> <cluster-ids>
,→<attribute-ids> <attribute-value>

Note:
• attribute_value should utilize a JSON object string. And the format of this string is the same as the command_data in cluster commands. This JSON object should contain only one item that represents the attribute
value.
Here are some examples of the attribute_value format.
For StartUpOnOff attribute of OnOff Cluster, you should use the following JSON structures as the attribute_value
to represent the StartUpOnOff 2 and null:
matter esp controller write-attr <node_id> <endpoint_id> 6 0x4003 "{\"0:U8\
,→": 2}"
matter esp controller write-attr <node_id> <endpoint_id> 6 0x4003 "{\"0:NULL\
,→": null}"

For Binding attribute of Binding cluster, you should use the following JSON structure as the attribute_value to
represent the binding list [{"node":1, "endpoint":1, "cluster":6}]:
matter esp controller write-attr <node_id> <endpoint_id> 30 0 "{\"0:ARR-OBJ\
,→":[{\"1:U64\":1, \"3:U16\":1, \"4:U32\": 6}]}"

For ACL attribute of AccessControl cluster, you should use the following JSON structure as the attribute_value
to represent the AccessControlList [{"privilege": 5, "authMode": 2, "subjects": [112233],
"targets": null}, {"privilege": 4, "authMode": 3, "subjects": [1], "targets":
null}]:
matter esp controller write-attr <node_id> <endpoint_id> 31 0 "{\"0:ARR-OBJ\
,→":[{\"1:U8\": 5, \"2:U8\": 2, \"3:ARR-U64\": [112233], \"4:NULL\": null},
,→{\"1:U8\": 4, \"2:U8\": 3, \"3:ARR-U64\": [1], \"4:NULL\": null}]}"

To write multiple attributes in one commands, the attribute_value should be a JSON array. For example, to write
the ACL attribute and Binding attribute above, you should use the following JSON structure as the attribute_value:
matter esp controller write-attr <node_id> <endpoint_id1>,<endpoint_id2> 31,
,→30 0,0 "[{\"0:ARR-OBJ\":[{\"1:U8\": 5, \"2:U8\": 2, \"3:ARR-U64\":␣
,→[112233], \"4:NULL\": null}, {\"1:U8\": 4, \"2:U8\": 3, \"3:ARR-U64\": [1],
,→ \"4:NULL\": null}]}, {\"0:ARR-OBJ\":[{\"1:U64\":1, \"3:U16\":1, \"4:U32\
,→": 6}]}]"

For attributes of type uint64_t or int64_t, if the absolute value is greater than (2^53), you should use string to represent
number in JSON structure for precision
matter esp controller write-attr <node_id> <endpoint_id> 42 0 "{\"0:ARR-OBJ\
,→":[{\"1:U64\": \"9007199254740993\", \"2:U8\": 0}]}"

1.2. 2. Developing with the SDK

35


```

