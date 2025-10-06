# Page 49

## Text Content

```
ESP-Matter Programming Guide, Release latest

3.5.1 Setup Thread BR
The otbr-posix can be run on RaspberryPi or Ubuntu machine. Connecting an RCP to the host, the port RCP_PORT1
for it will be /dev/ttyUSBX or /dev/ttyACMX.
• Build the otbr-posix on the host
git clone https://github.com/openthread/ot-br-posix
cd ot-br-posix
./script/bootstrap
./script/setup

Then the otbr-posix will be built and a service named otbr-agent will be created on the host. You can disable the service
and start the otbr-posix manually.
sudo systemctl disable otbr-agent.service
sudo ./build/otbr/src/agent/otbr-agent -I wpan0 -B eth0 -v spinel+hdlc+uart://{RCP_
,→PORT1}

In the above commands:
• wpan0 is the infra network interface. The network interface named wpan0 will be created on the host as the thread
network interface.
• eth0 is the backbone network interface, which is always the ethernet or wifi network interface on the host, please
ensure that the backbone network interface is connected to the AP which the Wi-Fi product is also connected to.
• RCP_PORT1 is the port of RCP for Thread BR.
The otbr-posix is running on the host now. Open another terminal, start console for otbr-posix, form Thread network,
and get dataset.
sudo ot-ctl
> ifconfig up
> thread start
> dataset active -x

Please record the dataset you get with the last command, it will be used by otcli-posix to join the BR’s network in the
next step.
3.5.2 Setup Thread End Device
We use the Posix Thread Command-Line Interface (CLI) as the Thread End Device. Connect another RCP to the host
and get the port RCP_PORT2 for it.
• Build the otcli on the host
git clone --recursive https://github.com/openthread/openthread.git
cd openthread/
./script/bootstrap
./bootstrap
./script/cmake-build posix
./build/posix/src/posix/ot-cli 'spinel+hdlc+uart:///dev/{RCP_PORT2}?uart,→baudrate=115200' -v

The console for the ot-cli will be started. Connect the ot-cli to the otbr’s Thread network with the dataset you got in the
above step.

1.3. 3. Matter Certification

45


```

