# Page 13

## Text Content

```
TACHAMMER LMR - D-10L20MP1-7

PRELIMINARY SPECIFICATIONS. SUBJECT TO CHANGE

3.1.3 LRA Drive Signal compatibility
TacHammer can be driven using LRA signals. While compatible, playback using LRA signals may result
in slightly reduced performance. For best results, shift the LRA drive signal frequency closer to the
TacHammer’s f0 (which is typically a lower frequency) or use audio waveforms. Using a closed-loop LRA
driver IC can help modify the LRA input signal to be optimized for higher output on the TacHammer.

3.1.4 Arbitrary and Audio waveform
TacHammer can be used to playback audio waveforms typically used for driving speakers. For best
results, signal amplification can be used via dedicated haptic ICs, audio ICs (PWM, Class-D, or Analog),
or discrete components (h-bridge).

3.1.5 Drive Signal conditioning
Low Pass Filtering can be used to eliminate audible sound generation. LPF can be accomplished using
ICs, discrete components, or software signal processing (DSP).
Resonance shifting can be achieved by applying a DC offset signal to the input waveform. A positive
offset will LOWER the unit resonance, and a negative offset will raise the unit resonance. WARNING:
using a DC offset will accelerate temperature rise in the coil. Be careful not to overheat the coil or the unit
may prematurely fail.

3.1.6 Non-Linear Magnetic Suspension Design
The TacHammer has a non-linear force profile, which is one of the primary differences between the
TacHammer and existing haptic technologies. This nonlinearity allows the TacHammer to have a wide
operating frequency range and quick response time primarily due to its low stiction at rest.
This means that you can create effective click effects with very low power consumption, respond well to
audio drive signals, and create very subtle haptic effects that would otherwise be difficult to produce on
legacy technologies. Additionally, the nonlinearity removes the requirement for expensive drive ICs,
making the overall solution cost lower than other HD haptic motors.

TITAN HAPTICS Inc. Patented US9716423B1
TITANHAPTICS.COM

|

CONFIDENTIAL AND PROPRIETARY

13/18


```

## Images

![Image from page 13](images/page_13_img_001.ppm)

![Image from page 13](images/page_13_img_002.ppm)

![Image from page 13](images/page_13_img_003.ppm)

