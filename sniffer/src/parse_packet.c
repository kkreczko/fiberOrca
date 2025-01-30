#include "parse_packet.h"
#include "common.h"

char* parse_packet(const struct timeval *tv, const u_char *payload) {
    static char packet_str[MAX_PACKET_STRING];
    char srcip[16], dstip[16];

    const struct ip *iphdr = (const struct ip *)payload;

    inet_ntop(AF_INET, &(iphdr->ip_src), srcip, sizeof(srcip));
    inet_ntop(AF_INET, &(iphdr->ip_dst), dstip, sizeof(dstip));

    const u_char *transport_payload = payload + (4 * iphdr->ip_hl);

    char timestamp[32];
    snprintf(timestamp, sizeof(timestamp), "%ld.%06ld", tv->tv_sec, tv->tv_usec);

    int src_port = 0, dst_port = 0;
    const char *proto_str = "UNKNOWN";

    switch (iphdr->ip_p) {
        case IPPROTO_TCP: {
            const struct tcphdr *tcphdr = (struct tcphdr*)transport_payload;
            proto_str = "TCP";
            src_port = ntohs(tcphdr->th_sport);
            dst_port = ntohs(tcphdr->th_dport);
            break;
        }

        case IPPROTO_UDP: {
            const struct udphdr *udphdr = (struct udphdr*)transport_payload;
            proto_str = "UDP";
            src_port = ntohs(udphdr->uh_sport);
            dst_port = ntohs(udphdr->uh_dport);
            break;
        }

        case IPPROTO_ICMP: {
            proto_str = "ICMP";
            break;
        }
    }

    snprintf(packet_str, MAX_PACKET_STRING,
            "%s;%d;%s;%s;%d;%s;%d;\n",
            proto_str,
            src_port,
            srcip,
            dstip,
            dst_port,
            timestamp,
            iphdr->ip_ttl);

    return packet_str;
}
