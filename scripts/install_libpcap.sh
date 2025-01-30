# You need to have make, gcc, bison and flex installed

curl -O https://www.tcpdump.org/release/libpcap-1.10.5.tar.xz
tar -xavf libpcap-1.10.5.tar.xz
cd libpcap-1.10.5
./configure
make
sudo make install