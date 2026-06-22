# fiberOrca

Real-time network packet analyzer.

## Features

- Real-time packet capture using libpcap
- Support for TCP, UDP, and ICMP protocols
- Detailed packet analysis including timestamps, ports, and IP addresses
- Live packet data visualization

## Requirements

- Linux
- root
- cmake
- clang

## Installation

1. Clone the repository:
```bash
git clone https://github.com/kruczys/fiberOrca.git
cd fiberOrca
```

2. Configure, build and install:
```bash
mkdir build && cd build
cmake ../ -DCMAKE_BUILD_TYPE=Relese
cmake --build .
sudo cmake --install .
```

## Running

```bash
fiberorca
```

Use ```--help``` for available start commands.

## Navigating the program

To navigate this program, use the following keybinds:

### Packet list
`q` - quit program

`f` - enter filter view

`r` - reset filter

`enter` - enter packet view

### Packet view
`q` - go back to packet list

### Filter 
`esc` - quit filter view without applying

`enter` - move to the next text input

## License

This project is licensed under the MIT License - see the LICENSE file for details.
