#ifndef INCLUDE_PACKET_HANDLER_H
#define INCLUDE_PACKET_HANDLER_H

#include "common.h"

void packet_handler(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr);
void packet_handler_ICP(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr);

#endif // INCLUDE_PACKET_HANDLER_H
