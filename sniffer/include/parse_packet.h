#ifndef INCLUDE_PARSE_PACKET_STRUCT_H
#define INCLUDE_PARSE_PACKET_STRUCT_H

#include <pcap.h>

#define MAX_PACKET_STRING 2048
#define PROTO_DATA_SIZE 1024

char* parse_packet(const struct pcap_pkthdr* header, const char* payload);

#endif //INCLUDE_PARSE_PACKET_STRUCT_H
