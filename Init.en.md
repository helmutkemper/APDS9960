# Configuration

## Wiring

This block initializes the I2C connection on the APDS9960 board and is required for I2C setup.

![](apds9960-arduino-wiring.svg)

<!-- place_the_control_panel_here -->

> A ratio near 1.0 on a single channel means that colour dominates the scene.

## Pinout

| Pin | Description                                                                                                                                                                                                                                                                         |
|-----|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Vin | this is the power pin. Since the sensor uses 3.3V, we have included an onboard voltage regulator that will take 3-5VDC and safely convert it down. To power the board, give it the same power as the logic level of your microcontroller - e.g. for a 5V micro like Arduino, use 5V |
| 3Vo | this is the 3.3V output from the voltage regulator, you can grab up to 100mA from this if you like                                                                                                                                                                                  |
| GND | common ground for power and logic                                                                                                                                                                                                                                                   |
| SCL | this is the I2C clock pin, connect to your microcontrollers I2C clock line. There is a 10K pullup on this pin and it is level shifted so you can use 3 - 5VDC.                                                                                                                      |
| SDA | this is the I2C data pin, connect to your microcontrollers I2C data line. There is a 10K pullup on this pin and it is level shifted so you can use 3 - 5VDC.                                                                                                                        |
| INT | this is the interrupt-output pin. It is 3V logic and you can use it to detect when a new reading is ready or when a reading gets too high or too low.                                                                                                                               |

## Example

Click on the image below to load the example.

![example image](/examples/example.png)

> This is an image with steganography; it hides binary project data. Do not edit it.
