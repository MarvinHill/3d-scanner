# 3d-scanner
This project contains software and 3d files for a 3d Scanner developed for the Virtual Reality Course at Hochschule Heilbronn in SoSE 2024.
## Structure
- configuration: contains ansible playbooks to configure the rpi and install the scanner firmware, and caddy webserver.
- firmware-rpi: contains firmware to controll the scanner
- frontent-scanner: contains the frontend for the webservice to controll the scanner with a graphical ui
- 3d: contains 3d files to build the scanner
## 3D Modells
## Used Hardware
- Raspberry Pi 1b
- m3 screws & m3 threat inserts
- eleego cabels
- ELEGOO stepper motor x3 and driver board ULN2003 x3
- 3d Printer & Filament
- endswitch
- power distibution hardware
### Pins
- motor one: pins           GPIO [8, 9, 7, 0]          ind [3 5 7 11]
- motor two: pins           GPIO [2, 3, 12, 13]        ind [13 15 19 21]
- motor three: pins         GPIO [15, 16, 1, 4]        ind [8 10 12 16]
- endswitch one: pins       GPIO [11, 10]              ind [26 24]
- endswitch two: pins       GPIO [6, 5]                ind [22 18]

### Driver
The gobot driver for stepper motors was used to controll the scanner 2 axis.

## Software
###
Linux PI Setup:

install git
install pi-blaster (for gobot)

### Gobot setup
[issue link](https://github.com/hybridgroup/gobot/issues/691)
### Build go
[How To Build](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

linux: 
command: GOOS=linux GOARCH=arm GOARM=6 go build -o ./build/scanner

windows:
(set GOOS=linux) & (set GOARCH=arm) & (set GOARM=6) & go build -o ./build/scanner

### Copy to rpi
scp scanner marvin@192.168.188.59:/tmp

### Development
## Result
