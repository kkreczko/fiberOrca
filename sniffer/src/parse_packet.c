#include "parse_packet.h"
#include "common.h"

char* parse_packet(const struct timeval *tv, const u_char *payload) {
    static char packet_str[MAX_PACKET_STRING];
    char proto_specific[512] = "";
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

            char flags[8];
            snprintf(flags, sizeof(flags),
                    "%c%c%c%c%c%c",
                    (tcphdr->th_flags & TH_URG ? 'U' : '*'),
                    (tcphdr->th_flags & TH_ACK ? 'A' : '*'),
                    (tcphdr->th_flags & TH_PUSH ? 'P' : '*'),
                    (tcphdr->th_flags & TH_RST ? 'R' : '*'),
                    (tcphdr->th_flags & TH_SYN ? 'S' : '*'),
                    (tcphdr->th_flags & TH_FIN ? 'F' : '*'));

            snprintf(proto_specific, sizeof(proto_specific),
                    "FLAGS:%s;SEQ:%u;ACK:%u;WIN:%u;TCPLEN:%d",
                    flags,
                    ntohl(tcphdr->th_seq),
                    ntohl(tcphdr->th_ack),
                    ntohs(tcphdr->th_win),
                    4*tcphdr->th_off);
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
            const struct icmp *icmphdr = (struct icmp*)transport_payload;
            proto_str = "ICMP";
            snprintf(proto_specific, sizeof(proto_specific),
                    "TYPE:%d;CODE:%d;ID:%d;SEQ:%d",
                    icmphdr->icmp_type,
                    icmphdr->icmp_code,
                    ntohs(icmphdr->icmp_hun.ih_idseq.icd_id),
                    ntohs(icmphdr->icmp_hun.ih_idseq.icd_seq));
            break;
        }
    }

    snprintf(packet_str, MAX_PACKET_STRING,
            "%s;%d;%s;%s;%d;%s;ID:%d;TOS:%d;TTL:%d;IPLEN:%d;%s",
            proto_str,
            src_port,
            srcip,
            dstip,
            dst_port,
            timestamp,
            ntohs(iphdr->ip_id),
            iphdr->ip_tos,
            iphdr->ip_ttl,
            4*iphdr->ip_hl,
            proto_specific);

    return packet_str;
}
