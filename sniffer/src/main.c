#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <pcap/pcap.h>
#include <netinet/tcp.h>
#include <netinet/udp.h>
#include <netinet/ip_icmp.h>

pcap_t* handle;
int linkhdrlen;
int packets;
pcap_handler packet_handler;

pcap_t* create_pcap_handle(char* device, char* filter);
int get_link_header_len(pcap_t *handle);
void stop_capture(int sig_number);

int main(int argc, char* argv[]) {
    char device[256];
    char filter[256];
    int count = 0;
    int opt;

    *device = 0;
    *filter = 0;

    while ((opt = getopt(argc, argv, "hi:n:") != -1)) {
        switch (opt) {
            case 'h':
                printf("you need some help? waawaa go cry about it %s", argv[0]);
                EXIT_SUCCESS;
                break;
            case 'i':
                strcpy(device, optarg);
                break;
            case 'n':
                count = atoi(optarg);
                break;
        }
    }
    
    for (int i = optind; i < argc; i++) {
        strcat(filter, argv[i]);
        strcat(filter, " ");
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
        return EXIT_FAILURE;
    }
    
    if (pcap_loop(handle, count, packet_handler, (u_char*)NULL) < 0) {
        fprintf(stderr, "pcap_loop failed wawa: %s\n", pcap_geterr(handle));
        return EXIT_FAILURE;
    }

    stop_capture(0);
}
