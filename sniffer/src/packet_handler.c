#include <packet_handler.h>
#include <connect_and_send.h>
#include <parse_packet.h>

#include <arpa/inet.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

//WARN When switching to ICP this will have to be changed, or we should add new function that will work only with ICP
//we might leave non-ICP functionality of this program for UX
void packet_handler(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr) {
    struct icmp *icmphdr;
    struct tcphdr *tcphdr;
    struct udphdr *udphdr;

    char iphdrInfo[256];
    char srcip[256];
    char dstip[256];

    packetptr += linkhdrlen;
    struct ip *iphdr = (struct ip *) packetptr;

    if (iphdr->ip_v != 4) {
        printf("Not an IPv4 packet. Skipping...\n\n");
        return;
    }

    inet_ntop(AF_INET, &(iphdr->ip_src), srcip, sizeof(srcip));
    inet_ntop(AF_INET, &(iphdr->ip_dst), dstip, sizeof(dstip));

    snprintf(iphdrInfo, sizeof(iphdrInfo), "ID:%d TOS:0x%x, TTL:%d IpLen:%d DgLen:%d",
            ntohs(iphdr->ip_id), iphdr->ip_tos, iphdr->ip_ttl,
            4*iphdr->ip_hl, ntohs(iphdr->ip_len));

    packetptr += 4*iphdr->ip_hl;

    switch (iphdr->ip_p) {
        case IPPROTO_TCP:
            tcphdr = (struct tcphdr*)packetptr;
            printf("TCP  %s:%d -> %s:%d\n", srcip, ntohs(tcphdr->th_sport),
                   dstip, ntohs(tcphdr->th_dport));
            printf("%s\n", iphdrInfo);
            printf("%c%c%c%c%c%c Seq: 0x%x Ack: 0x%x Win: 0x%x TcpLen: %d\n",
                   (tcphdr->th_flags & TH_URG ? 'U' : '*'),
                   (tcphdr->th_flags & TH_ACK ? 'A' : '*'),
                   (tcphdr->th_flags & TH_PUSH ? 'P' : '*'),
                   (tcphdr->th_flags & TH_RST ? 'R' : '*'),
                   (tcphdr->th_flags & TH_SYN ? 'S' : '*'),
                   (tcphdr->th_flags & TH_FIN ? 'F' : '*'),
                   ntohl(tcphdr->th_seq), ntohl(tcphdr->th_ack),
                   ntohs(tcphdr->th_win), 4*tcphdr->th_off);
            break;

        case IPPROTO_UDP:
            udphdr = (struct udphdr*)packetptr;
            printf("UDP  %s:%d -> %s:%d\n", srcip, ntohs(udphdr->uh_sport),
                   dstip, ntohs(udphdr->uh_dport));
            printf("%s\n", iphdrInfo);
            break;

        case IPPROTO_ICMP:
            icmphdr = (struct icmp*)packetptr;
            printf("ICMP %s -> %s\n", srcip, dstip);
            printf("%s\n", iphdrInfo);
            printf("Type:%d Code:%d ID:%d Seq:%d\n", icmphdr->icmp_type, icmphdr->icmp_code,
                   ntohs(icmphdr->icmp_hun.ih_idseq.icd_id),
                   ntohs(icmphdr->icmp_hun.ih_idseq.icd_seq));
            break;

        default:
            printf("Unknown protocol: %d\n", iphdr->ip_p);
            break;
    }

    printf("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n\n");
    packets++;
}

void packet_handler_IPC(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr) {
    static int sock = -1;

    if (sock == -1) {
        sock = create_socket();
        if (sock == -1) {
            fprintf(stderr, "Could not create socket\n");
            return;
        }
    }

    packetptr += linkhdrlen;

    if (((struct ip *)packetptr)->ip_v != 4) {
        fprintf(stderr, "Not an IPv4 packet. Skipping...\n\n");
        return;
    }

    char *packet_str = parse_packet(&packethdr->ts, packetptr);
    if (send_data(sock, packet_str, strlen(packet_str)) == -1) {
        close(sock);
        sock = -1;
    }
}