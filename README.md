# APDS9960

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

## Description

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
