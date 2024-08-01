# 3d-scanner
This project contains software and 3d files for a 3d Scanner developed for the Virtual Reality Course at Hochschule Heilbronn in SoSE 2024.
## Showcase

![blueprint](https://github.com/user-attachments/assets/81f2d787-ee72-472c-a32e-a927f29f6c98)

https://github.com/user-attachments/assets/66f47e95-8ec8-498e-958c-c9a36b86d276


## Features
- Manual control of the scanner movement on two axis
- Motor leveling of the axis that holds the camera with end-switches
- Automated photo mode where the scanner takes a specified amount of photos from multiple distributed angles, so that the images cover the object from all Sites
## Structure
- configuration: contains ansible playbooks to configure the rpi and install the scanner firmware, and caddy webserver.
- firmware-rpi: contains firmware to controll the scanner
- frontent-scanner: contains the frontend for the webservice to controll the scanner with a graphical ui
- 3d: contains 3d files to build the scanner
## 3D Modells
stl files can be found in the 3d folder, look at the blueprint in the readme to understand where which part fits
## Used Hardware
- Raspberry Pi zero 2w
- m3 screws & m3 threat inserts
- eleego cabels
- ELEGOO stepper motor x3 and driver board ULN2003 x3
- 3d Printer & Filament
- endswitch
### Pins
- motor table axis: pins           index [3 5 7 11]
- motor one camera: pins           index [23 29 31 33]
- motor two camera: pins           index [8 10 12 16]
- endswitch one: pins              index [in:36, out:35]
- endswitch two: pins              index [in:38, out:37]

### Driver
The gobot driver for stepper motors was used to controll the scanner 2 axis.

## Software
### Linux RPI Configuration

install ansible and docker on host system
install pi-blaster (for gobot if older rpi is used, not needed for newer) 

### Gobot setup
oder version of rpi
[issue link](https://github.com/hybridgroup/gobot/issues/691)
### Build go
[How To Build](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

linux: 
command: GOOS=linux GOARCH=arm GOARM=6 go build -o ./build/scanner

windows:
(set GOOS=linux) & (set GOARCH=arm) & (set GOARM=6) & go build -o ./build/scanner


- run docker compose to build the artifacts
- then run the ansible playbook `ansible-playbook -i production site.yml`

### Copy to rpi
scp scanner marvin@192.168.188.59:/tmp
