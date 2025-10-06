# Page 9

## Text Content

```
titanhaptics.com

Command Syntax can be referred to in the following chart:
Channel
(CHNL)

CHANNEL
0=L, R and M, 1=L, 2=R, 3=M

Usage: Determines the channel for generating the output. 0=L, R and M, 1=L, 2=R, 3=M
Example: CHNL 3; Tick 0.85 20.5;
Result: Single Tick will be generated from Channel 3 (M channel)
Pulse

STRENGTH
From 0 to 1

DURATION
in milliseconds

Usage: A single soft impulse for pulses, bumps, wobbles and motion forces
Example: CHNL 2; Pulse 0.85 20;
Result: Single Pulse at 85% strength, for 20 ms, output to 2 channel (R channel).
Tick

STRENGTH
From 0 to 1

DURATION
in milliseconds

Usage: A single sharp impulse for generating clicks, textures, taps and impacts.
Example: CHNL 0; Tick 0.85 20;
Result: Single tick at 85% strength, for 20 ms, output to All channel (L, R and M channel)
Pause

DURATION
in milliseconds

Usage: Everything is paused for certain duration.
Example: CHNL 1; vibrate 100 0.3 15000 1 1;
Result: Pause of 20 ms between two Tick command while the is generated from Channel 1
Vibrate

FREQUENCY
in hertz

STRENGTH
From 0 to 1

DURATION
in milliseconds

DUTY CYCLE
from 0 to 1

SHARPNESS
from 0 to 1

Usage: A full swing vibration used to create alerts, textures etc. Use ‘Duty Cycle’ to vary the
‘fullness’ of the vibration and ‘Sharpness’ to vary the waveform between pure sine (0,
rounded) and pure square (1, very sharp).
Example: CHNL 1;vibrate 100 0.3 15000 1 1;
Result: Vibration at 120 HZ, 30% strength, for 15000 ms, at 100% duty cycle and 100%
sharpness.
PCM

Frame Frequency
Hz

Frame Size

PCM value
8 bit value (from 0 to 255)

Usage: A PCM (Pulse-Code-Modulated) is a common format for audio based haptics. It is
used to convert 8 bit audio data to vibration.
Example: CHNL 0;F 4000 20;PCM 244,0,0,0,0,0,0,0,0,0,0,255,0,0,0,0,45,45,45,45;
Result: Stream of 8 bit PCM data is generated with frame frequency 4000 and frame size
20 in channel 0.

PRIVATE AND CONFIDENTIAL. Property of TITAN HAPTICS INC. All Rights Reserved. www.titanhaptics.com

9 / 13


```

## Images

![Image from page 9](images/page_9_img_001.jpg)

![Image from page 9](images/page_9_img_002.ppm)

