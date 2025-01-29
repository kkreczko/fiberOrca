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
            perror("pcap_findalldevs");
            return NULL;
        }
        if (devices == NULL) {
            perror("pcap_findalldevs");
            return NULL;
        }
        strncpy(device, devices[0].name, 255);
        device[255] = '\0';
        pcap_freealldevs(devices);
    }

    if (pcap_lookupnet(device, &srcip, &netmask, errbuf) == PCAP_ERROR) {
        perror("pcap_lookupnet");
        netmask = 0;
        srcip = 0;
    }

    handle = pcap_open_live(device, BUFSIZ, 1, 1000, errbuf);
    if (handle == NULL) {
        perror("pcap_open_live");
        return NULL;
    }

    if (filter[0] != '\0') {
        if (pcap_compile(handle, &bpf, filter, 0, netmask) == PCAP_ERROR) {
            perror("pcap_compile");
            pcap_close(handle);
            return NULL;
        }

        if (pcap_setfilter(handle, &bpf) == PCAP_ERROR) {
            perror("pcap_setfilter");
            pcap_close(handle);
            return NULL;
        }

        pcap_freecode(&bpf);
    }

    return handle;
}

//WARN we can add functionality to catch Wifi packets with their cool information like signal power etc.
//right now we can catch only those wifi packets which were converted to Ethernet by network interface card driver
//IEEE 802.11
int get_link_header_len(pcap_t *handle) {
    int linktype;

    if ((linktype = pcap_datalink(handle)) == PCAP_ERROR) {
        perror("pcap_datalink");
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

void stop_capture() {
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

void stop_capture_IPC() {
    if (handle)
        pcap_close(handle);

    exit(EXIT_SUCCESS);
}