#ifndef INCLUDE_PACKET_HANDLER_H
#define INCLUDE_PACKET_HANDLER_H

#include "common.h"

void PacketHandlerVerbose(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr);
void PacketHandlerIPC(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr);
void PacketHandlerTest(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr);

#endif // INCLUDE_PACKET_HANDLER_H
