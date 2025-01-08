#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <pcap/pcap.h>
#include <netinet/tcp.h>
#include <netinet/udp.h>
#include <netinet/ip_icmp.h>
#include <netinet/ip.h>
#include <arpa/inet.h>

pcap_t* handle = NULL;
int linkhdrlen = 0;
int packets = 0;

//TODO move methods to their own header and source files
//TODO switch from printing to ICP communication with bubbletea process
void stop_capture(int sig_number) {
    struct pcap_stat stats;

    if (handle && pcap_stats(handle, &stats) >= 0) {
        printf("\n%d packets captured\n", packets);
        printf("%d packets received\n", stats.ps_recv);
        printf("%d packets dropped\n\n", stats.ps_drop);
    }

    if (handle) {
        pcap_close(handle);
    }
    exit(EXIT_SUCCESS);
}

void packet_handler(u_char *user, const struct pcap_pkthdr *packethdr, const u_char *packetptr) {
    struct ip *iphdr;
    struct icmp *icmphdr;
    struct tcphdr *tcphdr;
    struct udphdr *udphdr;

    char iphdrInfo[256];
    char srcip[256];
    char dstip[256];

    packetptr += linkhdrlen;
    iphdr = (struct ip*)packetptr;

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

int get_link_header_len(pcap_t *handle) {
    int linktype;

    if ((linktype = pcap_datalink(handle)) == PCAP_ERROR) {
        fprintf(stderr, "pcap_datalink(): %s\n", pcap_geterr(handle));
        return 0;
    }

    switch (linktype) {
        case DLT_NULL:
            return 4;
        case DLT_EN10MB:
            return 14;
        case DLT_SLIP:
        case DLT_PPP:
            return 24;
        default:
            printf("Unsupported data link (%d)\n", linktype);
            return 0;
    }
}

pcap_t* create_pcap_handle(char* device, char* filter) {
    char errbuf[PCAP_ERRBUF_SIZE];
    pcap_if_t *devices = NULL;
    struct bpf_program bpf;
    bpf_u_int32 netmask;
    bpf_u_int32 srcip;

    if (!device[0]) {
        if (pcap_findalldevs(&devices, errbuf) == PCAP_ERROR) {
            fprintf(stderr, "pcap_findalldevs(): %s\n", errbuf);
            return NULL;
        }
        if (devices == NULL) {
            fprintf(stderr, "No devices found\n");
            return NULL;
        }
        strncpy(device, devices[0].name, 255);
        device[255] = '\0';
        pcap_freealldevs(devices);
    }

    if (pcap_lookupnet(device, &srcip, &netmask, errbuf) == PCAP_ERROR) {
        fprintf(stderr, "pcap_lookupnet(): %s\n", errbuf);
        netmask = 0;
        srcip = 0;
    }

    handle = pcap_open_live(device, BUFSIZ, 1, 1000, errbuf);
    if (handle == NULL) {
        fprintf(stderr, "pcap_open_live(): %s\n", errbuf);
        return NULL;
    }

    if (filter[0] != '\0') {
        if (pcap_compile(handle, &bpf, filter, 0, netmask) == PCAP_ERROR) {
            fprintf(stderr, "pcap_compile(): %s\n", pcap_geterr(handle));
            pcap_close(handle);
            return NULL;
        }

        if (pcap_setfilter(handle, &bpf) == PCAP_ERROR) {
            fprintf(stderr, "pcap_setfilter(): %s\n", pcap_geterr(handle));
            pcap_close(handle);
            return NULL;
        }

        pcap_freecode(&bpf);
    }

    return handle;
}

int main(int argc, char* argv[]) {
    char device[256] = {0};
    char filter[256] = {0};
    int count = 0;
    int opt;

    while ((opt = getopt(argc, argv, "hi:n:")) != -1) {
        switch (opt) {
            case 'h':
                printf("Usage: %s [-i interface] [-n count] [filter]\n", argv[0]);
                return EXIT_SUCCESS;
            case 'i':
                strncpy(device, optarg, sizeof(device) - 1);
                break;
            case 'n':
                count = atoi(optarg);
                break;
            default:
                fprintf(stderr, "Invalid option\n");
                return EXIT_FAILURE;
        }
    }

    for (int i = optind; i < argc; i++) {
        if (strlen(filter) + strlen(argv[i]) + 2 <= sizeof(filter)) {
            strcat(filter, argv[i]);
            strcat(filter, " ");
        }
    }

    signal(SIGINT, stop_capture);
    signal(SIGTERM, stop_capture);
    signal(SIGQUIT, stop_capture);

    handle = create_pcap_handle(device, filter);
    if (handle == NULL) {
        return EXIT_FAILURE;
    }

    linkhdrlen = get_link_header_len(handle);
    if (linkhdrlen == 0) {
        pcap_close(handle);
        return EXIT_FAILURE;
    }

    if (pcap_loop(handle, count, packet_handler, (u_char*)NULL) < 0) {
        fprintf(stderr, "pcap_loop failed: %s\n", pcap_geterr(handle));
        pcap_close(handle);
        return EXIT_FAILURE;
    }

    stop_capture(0);
    return EXIT_SUCCESS;
}