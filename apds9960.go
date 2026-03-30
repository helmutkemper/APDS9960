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
// icon:sun
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

