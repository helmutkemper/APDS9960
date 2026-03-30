// Package blackbox
//
// APDS9960 is a colour, proximity, and gesture sensor connected via I2C.
//
// Place Init at the global scope. Place Run inside a Loop.
// Wire I2CBus.bus → APDS9960.i2c before generating code.
package blackbox

import (
	"machine"
	"time"
)

// APDS9960 reads colour (RGBC) data via I2C.
//
// icon:lightbulb
type APDS9960 struct {
	i2c   *machine.I2C
	gain  byte `prop:"ADC Gain"         default:"0"   options:"0,1,2,3"`
	atime byte `prop:"Integration Time" default:"255"`
}

// Init configures the sensor on the given I2C bus.
//
// icon:hourglass-start.
//
// Params
//
//	i2c: I2C bus — wire from an I2CBus Init block.  connection:mandatory.  unit:i2c_bus.
//
// Returns
//
//	err: initialisation error.  connection:optional.
func (s *APDS9960) Init(i2c *machine.I2C) (err error) {
	s.i2c = i2c
	s.i2c.WriteRegister(0x39, 0x80, []byte{0x01})
	s.i2c.WriteRegister(0x39, 0x81, []byte{s.atime})
	s.i2c.WriteRegister(0x39, 0x8F, []byte{s.gain})
	s.i2c.WriteRegister(0x39, 0x80, []byte{0x03})
	return nil
}

// Run reads the four RGBC colour channels.
//
// executionOrder:10. icon:sun.
//
// Returns
//
//	clear: total light intensity.  range:0..65535.  unit:lux_counts.   connection:optional.
//	red:   red channel.            range:0..65535.  unit:color_counts.  connection:optional.
//	green: green channel.          range:0..65535.  unit:color_counts.  connection:optional.
//	blue:  blue channel.           range:0..65535.  unit:color_counts.  connection:optional.
func (s *APDS9960) Run() (clear, red, green, blue uint16) {
	data := make([]byte, 8)
	s.i2c.ReadRegister(0x39, 0x94, data)
	clear = uint16(data[0]) | uint16(data[1])<<8
	red = uint16(data[2]) | uint16(data[3])<<8
	green = uint16(data[4]) | uint16(data[5])<<8
	blue = uint16(data[6]) | uint16(data[7])<<8
	return
}

// Log reads the four RGBC colour channels.
//
// executionOrder:20. icon:usb.
//
// Params
//
//	clear: total light intensity.  range:0..65535.  unit:lux_counts.   connection:mandatory.
//	red:   red channel.            range:0..65535.  unit:color_counts.  connection:mandatory.
//	green: green channel.          range:0..65535.  unit:color_counts.  connection:mandatory.
//	blue:  blue channel.           range:0..65535.  unit:color_counts.  connection:mandatory.
func (s *APDS9960) Log(clear, red, green, blue uint16) {
	println("C:", clear, "R:", red, "G:", green, "B:", blue)
	time.Sleep(500 * time.Millisecond)
}

/*
manualName:Wiring-guide.
language:en.
showIn:init.
```markdown
# APDS9960 — Wiring Guide

| APDS9960 | Pico Pin | Notes      |
|----------|----------|------------|
| SDA      | GP4      | I2C data   |
| SCL      | GP5      | I2C clock  |
| VCC      | 3V3      | 3.3 V only |
| GND      | GND      |            |

Place an **I2CBus Init** block first and wire its **bus** output here.
```
*/

/*
manualName:Reading-colors.
language:en.
showIn:run.
```markdown
# Reading Colour Values

Divide each channel by `clear` to get a lighting-independent ratio:

```
redRatio   = red   / clear
greenRatio = green / clear
blueRatio  = blue  / clear
```

A ratio near 1.0 on a single channel means that colour dominates the scene.
```
*/

/*
manualName:Board.
language:en.
showIn:run.
```markdown
# APDS9960 Board

## Description of the board APDS9960

This handy sensor is full of features! Add basic gesture sensing, RGB color sensing, proximity sensing, or ambient
light sensing to your project with the Adafruit APDS9960 Proximity, Light, RGB and Gesture Sensor. When connected to
your microcontroller (running our library code) it can detect simple gestures (left to right, right to left, up to down,
down to up are currently supported), return the amount of red, blue, green, and clear light, or return how close an
object is to the front of the sensor. This device uses an I2C interface so it's easy to wire up and use.

The APDS9960 from Avago Technologies has an integrated IR LED and driver, along with four directional photodiodes that
sense reflected IR energy from the LED. It's proximity detection feature allows it to measure the distance an object is
from the front of the sensor (up to a few centimeters) with 8 bit resolution.

Since there are four IR sensors, you can measure the changes in light reflectance at each of the cardinal locations over
time and turn those changes into gestures. Our interface library can detect directional gestures (left to right, right
to left, up to down, down to up), but in theory more complicated gestures like zig-zag, clockwise or counterclockwise
circle, near to far, etc. could also be detected with additional code.

The APDS9960 has a configurable interrupt that can fire when a certain proximity threshold is broken, or when a color
sensor breaks a certain threshold.

For your convenience we've pick-and-placed the sensor on a PCB with a 3.3V regulator and some level shifting so it can
be easily used with your favorite 3.3V or 5V microcontroller.

A ratio near 1.0 on a single channel means that colour dominates the scene.
```
*/

/*
manualName:Pin out.
language:en.
showIn:run.
```markdown
# APDS9960 Board

## Pin out

```
 +------------------------+
 |        TOP VIEW        |
 |                    (◎) | INT
 |  +-------+----+    (◎) | SDA
 |  | ∩     |  ∩ |    (◎) | SCL
 |  | ∪     |  ∪ |    (◎) | GND
 |  +-------+----+    (◎) | Out 3.3V
 |                    (◎) | VIn 3.3V-5V
 |    COMPONENTS SIDE     |
 +------------------------+
```

 | Pin | Description                                                                                                                                                                                                                                                                         |
 |-----|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
 | INT | this is the interrupt-output pin. It is 3V logic and you can use it to detect when a new reading is ready or when a reading gets too high or too low.                                                                                                                               |
 | SDA | this is the I2C data pin, connect to your microcontrollers I2C data line. There is a 10K pullup on this pin and it is level shifted so you can use 3 - 5VDC.                                                                                                                        |
 | SCL | this is the I2C clock pin, connect to your microcontrollers I2C clock line. There is a 10K pullup on this pin and it is level shifted so you can use 3 - 5VDC.                                                                                                                      |
 | GND | common ground for power and logic                                                                                                                                                                                                                                                   |
 | OUT | this is the 3.3V output from the voltage regulator, you can grab up to 100mA from this if you like                                                                                                                                                                                  |
 | VIn | this is the power pin. Since the sensor uses 3.3V, we have included an onboard voltage regulator that will take 3-5VDC and safely convert it down. To power the board, give it the same power as the logic level of your microcontroller - e.g. for a 5V micro like Arduino, use 5V |

```*/

/*
manualName:Program guide.
language:en.
showIn:run.
```markdown
# APDS9960 Board

## Program guide

```
 +------------------------+             +------------------------+
 |        I2C INIT        |             |      APDS9960 INIT     |
 +------------------------+             +------------------------+
 |                i2c (◉)-+-------------+-(◉) i2c      error (◎) |
 +------------------------+             +------------------------+

 + >>>------------------------ loop ------------------------->>> +
 |                                                               |
 |  +------------------------+       +------------------------+  |
 |  |        I2C INIT        |       |       APDS9960 LOG     |  |
 |  +------------------------+       +------------------------+  |
 |  |              clear (◉)-+-------+-(◉) clear              |  |
 |  |                red (◉)-+-------+-(◉) red                |  |
 |  |              green (◉)-+-------+-(◉) green              |  |
 |  |               blue (◉)-+-------+-(◉) blue               |  |
 |  |                        |       |           serial / USB |  |
 |  +------------------------+       +------------------------+  |
 |                                                               |
 + <<<------------------------ loop -------------------------<<< +

```

Source: (adafruit)[https://learn.adafruit.com/adafruit-apds9960-breakout/]
```*/
