#ifndef INCLUDE_CONNECT_AND_SEND_H
#define INCLUDE_CONNECT_AND_SEND_H

#include <pcap.h>

int create_socket();
int send_packet(int sockfd, struct pcap_pkthdr* header);

#endif // INCLUDE_CONNECT_AND_SEND_H