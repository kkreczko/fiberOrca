# fiberOrca

Real-time network packet analyzer with frontend written in Go and backend in C. The application uses Unix Domain Sockets for IPC communication between components.

## Features

- Real-time packet capture using libpcap
- Support for TCP, UDP, and ICMP protocols
- Detailed packet analysis including timestamps, ports, and IP addresses
- Low-latency IPC communication between components
- Live packet data visualization

## Requirements

- Linux-based system
- libpcap
- root privileges

## Installation

1. Clone the repository:
```bash
git clone https://github.com/kruczys/fiberOrca.git
cd fiberOrca
```

2. Install libpcap dependencies (you need to have make, gcc, bison and flex installed for libpcap to install):
```bash
chmod +x scripts/install_libpcap.sh
sudo ./scripts/install_libpcap.sh
```

## Running from fiberOrca's directory

The project includes a run script that handles starting both components:

```bash
chmod +x scripts/run.sh
sudo ./scripts/run.sh
```

Out of the box functionality of running this program from anywhere within system will be added later. For now you can add alias of absolute path to run.sh to your .bashrc file. 

## License

This project is licensed under the MIT License - see the LICENSE file for details.
