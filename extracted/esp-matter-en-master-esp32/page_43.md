# Page 43

## Text Content

```
ESP-Matter Programming Guide, Release latest

2.11.2 Cluster Implementation
The custom cluster should be implemented after the app-common code has been generated.
2.11.2.1 Custom Cluster Attributes
The attributes in a cluster can be managed with the two following methods. A cluster can utilize both the methods to
manage its attributes.
• Attribute Accessors
By default, all the attributes are stored in the ZCL data model and can be managed with the Attribute Accessors
generated in the app-common code. You can set/get the attribute values with the Accessors APIs.
• Attribute Access Interface (AAI)
Matter provides a virtual class, AttributeAccessInterface, which can be inherited by the custom cluster to manage its attributes. Attributes managed by AAI should be added to attributeAccessInterfaceAttributes array in both the zcl configuration file and the zcl test configuration file. Then, run the
zap_regen_all.py to regenerate the app-common code. Once the code is regenerated, the Attribute Accessors APIs for attributes managed by AAI will be removed.
Notes that attributes of complex types(structure or array) cannot be handled by Attribute Accessors and MUST be
managed using AAI.
2.11.2.2 Custom Cluster Commands
The commands in a cluster can be handled with one of the two following methods. A cluster can only choose one method
to handle its commands.
• Ember Command callbacks
By default, all the commands are handled using Ember command callbacks. The zap tool generates declarations
for these callbacks in the app-common code. And the corresponding definitions should be implemented to use the
commands within the clusters.
• Command Handler Interface (CHI)
Matter also provides a virtual class, CommandHandlerInterface, which can be inherited to handle commands within the cluster. If the commands in a cluster are handled by CHI. The cluster should be added to the
CommandHandlerInterfaceOnlyClusters array in the zap configuration data file. After modifying the
zap configuration data, the code should be regenerated, which will remove the Ember command callback declarations.
2.11.2.3 Custom Cluster Events
All the events are managed by the EventLogging. If an event is triggered, chip::app::LogEvent() can be called
to record it. The event will then be reported to the subscriber that has subscribed to it.

1.2. 2. Developing with the SDK

39


```

