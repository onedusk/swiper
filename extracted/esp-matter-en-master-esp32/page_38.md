# Page 38

## Text Content

```
ESP-Matter Programming Guide, Release latest

matter esp controller invoke-cmd <node-id> <endpoint-id> 63 0 "{\"0:OBJ\": {\
,→"0:U16\": 42, \"1:U8\": 0, \"2:BYT\": \"0NHS09TV1tfY2drb3N3e3w==\", \"3:U64\":␣
,→2220000, \"4:NULL\": null, \"5:NULL\": null, \"6:NULL\": null, \"7:NULL\": null}
,→}"

• For AddGroup command in Groups cluster, the command-data ({"groupID":
"grp1"}) should be:

1, "groupName":

matter esp controller invoke-cmd <node-id> <endpoint-id> 0x4 0 "{\"0:U16\": 1, \
,→"1:STR\": \"grp1\"}"

2.10.4 Read commands
The read_command class is used for sending read commands to other end-devices. Its constructor function could
accept two callback inputs:
• Attribute report callback: This callback will be called upon the reception of the attribute report for read-attribute
commands.
• Event report callback: This callback will be called upon the reception of the event report for read-event commands.
2.10.4.1 Read attribute commands
The read-attr commands are used for sending the commands of reading attributes on end-devices.
• Send the read-attribute command:
matter esp controller read-attr <node-id> <endpoint-ids> <cluster-ids>
,→<attribute-ids>

Note:
• endpoint-ids can represent a single or multiple endpoints, e.g. ‘0’ or ‘0,1’. And the same applies to cluster-ids,
attribute-ids, and event-ids below.

2.10.4.2 Read event commands
The read-event commands are used for sending the commands of reading events on end-devices.
• Send the read-event command:
matter esp controller read-event <node-id> <endpoint-ids> <cluster-ids> <event,→ids>

34

Chapter 1. Table of Contents


```

