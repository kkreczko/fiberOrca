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

void stop_capture(int sig_number) {
    struct pcap_stat stats;
    
    // TODO Change "handle" later, it's here to silence LSP warnings because im LAZY!
    // WARN 
    if (pcap_stats(handle, &stats) >= 0) {
        printf("\n%d packets captured\n", packets);
        printf("%d packets received\n", stats.ps_recv);
        printf("%d packets dropped\n\n", stats.ps_drop);
    }

    pcap_close(handle);
    EXIT_SUCCESS;
}

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

    return EXIT_SUCCESS;
}
