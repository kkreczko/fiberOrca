#ifndef INCLUDE_COMMON_H
#define INCLUDE_COMMON_H

#include <pcap.h>
#include <pcap/pcap.h>
#include <netinet/ip.h>
#include <string.h>
#include <stdio.h>
#include <netinet/tcp.h>
#include <netinet/udp.h>
#include <netinet/ip_icmp.h>
#include <stdlib.h>
#include <arpa/inet.h>
#include <unistd.h>

extern pcap_t* handle;
extern int linkhdrlen;
extern u_int packets;

#endif // INCLUDE_COMMON_H
