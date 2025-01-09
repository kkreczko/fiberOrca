#include "capture.h"
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <arpa/inet.h>

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