# Wiring Guide

| APDS9960 | Pico Pin | Notes      |
|----------|----------|------------|
| SDA      | GP4      | I2C data   |
| SCL      | GP5      | I2C clock  |
| VCC      | 3V3      | 3.3 V only |
| GND      | GND      |            |

Place an **I2CBus Init** block first and wire its **bus** output here.
