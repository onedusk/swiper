# Page 40

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.10.6 Subscribe commands
The subscribe_command class is used for sending subscribe commands to other end-devices. Its constructor function
could accept four callback inputs:
• Attribute report callback: This callback will be invoked upon the reception of the attribute report for subscribeattribute commands.
• Event report callback: This callback will be invoked upon the reception of the event report for subscribe-event
commands.
• Subscribe done callback: This callback will be invoked when the subscription is terminated or shutdown.
• Subscribe failure callback: This callback will be invoked upon the failure of establishing CASE session.
2.10.6.1 Subscribe attribute commands
The subs-attr commands are used for sending the commands of subscribing attributes on end-devices.
• Send the subscribe-attribute command:
matter esp controller subs-attr <node-id> <endpoint-ids> <cluster-ids> <attribute,→ids> <min-interval> <max-interval>

2.10.6.2 Subscribe event commands
The subs-event commands are used for sending the commands of subscribing events on end-devices.
• Send the subscribe-event command:
matter esp controller subs-event <node-id> <endpoint-ids> <cluster-ids> <event,→ids> <min-interval> <max-interval>

2.10.7 Group settings commands
The group-settings commands are used to set group information of the controller. They are available when the
Enable matter commissioner option is enabled in menuconfig. If the controller wants to send multicast commands to end-devices, it should be in the same group as the end-devices.
• Set group information of the controller:
matter esp controller group-settings show-groups
matter esp controller group-settings add-group <group-id> <group-name>
matter esp controller group-settings remove-group <group-id>
matter esp controller group-settings show-keysets
matter esp controller group-settings add-keyset <ketset-id> <policy> <validity,→time> <epoch-key-oct-str>
matter esp controller group-settings remove-keyset <ketset-id>
matter esp controller group-settings bind-keyset <group-id> <ketset-id>
matter esp controller group-settings unbind-keyset <group-id> <ketset-id>

36

Chapter 1. Table of Contents


```

