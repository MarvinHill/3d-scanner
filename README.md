# 3d-scanner
This project contains software and 3d files for a 3d Scanner developed for the Virtual Reality Course at Hochschule Heilbronn in SoSE 2024.
## Structure
- configuration: contains ansible playbooks to configure the rpi and install the scanner firmware, and caddy webserver.
- firmware-rpi: contains firmware to controll the scanner
- frontent-scanner: contains the frontend for the webservice to controll the scanner with a graphical ui
- 3d: contains 3d files to build the scanner
## 3D Modells
## Used Hardware
- Raspberry Pi Zero 2W
- m3 screws & m3 threat inserts
- eleego cabels
- ELEGOO stepper motor x3 and driver board ULN2003 x3
- 3d Printer & Filament
- endswitch
- power distibution hardware
### Pins
- motor one: pins [8, 9, 7, 0]
- motor two: pins [2, 3, 12, 13]
- motor three: pins [15, 16, 1, 4]
- endswitch one: pins [11, 10]
- endswitch two: pins [6, 5]

### Driver
The gobot driver for stepper motors was used to controll the scanner 2 axis.

## Software
### Development
## Result
