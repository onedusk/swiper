# Page 42

## Text Content

```
ESP-Matter Programming Guide, Release latest

(continued from previous page)

<command source="client" code="0x01" name="CommandWithArgs" response=
,→"CommandWithArgsResponse" optional="false">
<description>
Command that takes two uint8 arguments.
</description>
<arg name="Arg1" type="int8u"/>
<arg name="Arg2" type="int8u"/>
</command>
<!-- Command Responses -->
<command source="server" code="0x02" name="CommandWithArgsResponse" optional=
,→"false" disableDefaultResponse="true">
<description>
Response for CommandwithArgs.
</description>
<arg name="ResponseArg" type="int8u"/>
</command>
<!-- Events -->
<event side="server" code="0x0000" name="TestEvent" priority="info"␣
,→isFabricSensitive="true" optional="false">
<description>
Example event with a event data
</description>
<field id="1" name="EventData" type="int32u"/>
</event>
</cluster>
</configurator>

The example XML file above illustrates a cluster with two attributes, two accepted commands, one generated command(command response), and one event.
After creating the custom cluster XML template file, add the root directory of your template file to the xmlRoot array
and the template file name to the xmlFile array in both the zcl configuration file and the zcl test configuration file.
Run zap_regen_all.py in Matter virtual environment to generate common code and client code for the custom
cluster.
cd esp_matter/connectedhomeip/connectedhomeip
source ./scripts/activate.sh
./scripts/tools/zap_regen_all.py

The code generation script will create client code for the custom cluster, supporting Android, Darwin, and Python controllers, as well as the chip-tool. It will also generate app-common code for the new custom cluster. The chip-tool can be
used to test the custom cluster after recompiling.

38

Chapter 1. Table of Contents


```

