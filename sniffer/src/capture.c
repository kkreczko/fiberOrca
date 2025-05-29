#include "capture.h"

pcap_t *CreatePcapHandle(char *device, char *filter) {
    char errBuff[PCAP_ERRBUF_SIZE];
    pcap_if_t *devices = NULL;
    struct bpf_program bpf;
    bpf_u_int32 netMask;
    bpf_u_int32 srcIp;

    if (!device[0]) {
        if (pcap_findalldevs(&devices, errBuff) == PCAP_ERROR) {
            perror("pcap_findalldevs\n");
            return NULL;
        }
        if (devices == NULL) {
            perror("pcap_findalldevs\n");
            return NULL;
        }
        strncpy(device, devices[0].name, 255);
        device[255] = '\0';
        pcap_freealldevs(devices);
    }

    if (pcap_lookupnet(device, &srcIp, &netMask, errBuff) == PCAP_ERROR) {
        perror("pcap_lookupnet\n");
        netMask = 0;
        srcIp = 0;
    }

    handle = pcap_open_live(device, BUFSIZ, 1, 1000, errBuff);
    if (handle == NULL) {
        perror("pcap_open_live\n");
        return NULL;
    }

    if (filter[0] != '\0') {
        if (pcap_compile(handle, &bpf, filter, 0, netMask) == PCAP_ERROR) {
            perror("pcap_compile\n");
            pcap_close(handle);
            return NULL;
        }

        if (pcap_setfilter(handle, &bpf) == PCAP_ERROR) {
            perror("pcap_setfilter\n");
            pcap_close(handle);
            return NULL;
        }

        pcap_freecode(&bpf);
    }

    return handle;
}

int GetLinkHeaderLen(pcap_t *handle) {
    int linkType;

    if ((linkType = pcap_datalink(handle)) == PCAP_ERROR) {
        perror("pcap_datalink\n");
        return 0;
    }

    switch (linkType) {
        case DLT_NULL:
            return 4;
        case DLT_EN10MB:
            return 14;
        case DLT_SLIP:
        case DLT_PPP:
            return 24;
        default:
            fprintf(stderr, "Error: Unknown link type encountered.\n");
            return 0;
    }
}

void StopCapture() {
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

int StopCaptureTest() {
    struct pcap_stat stats;

    if (handle && pcap_stats(handle, &stats) >= 0) {
        pcap_close(handle);
        return packets != packets - stats.ps_drop;
    }

    return 255;
}

void StopCaptureIPC() {
    if (handle)
        pcap_close(handle);

    exit(EXIT_SUCCESS);
}
