#include "capture.h"
#include "packet_handler.h"
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

enum mode {
    VERBOSE,
    TEST,
    IPC
};

int main(int argc, char* argv[]) {
    char device[256] = {0};
    char filter[256] = {0};
    int count = 0;
    int opt;
    int mode[3];
    mode[IPC] = 1;

    while ((opt = getopt(argc, argv, "hi:n:vt")) != -1) {
        switch (opt) {
            case 'h':
                printf("Usage: %s [-i interface] [-n count] [-v verbose] [-t test] [filter]\nexample: -n 64 tcp port 80", argv[0]);
                return EXIT_SUCCESS;
            case 'i':
                strncpy(device, optarg, sizeof(device) - 1);
                break;
            case 'n':
                count = atoi(optarg);
                break;
            case 'v':
                mode[VERBOSE] = 1;
                mode[IPC] = 0;
                break;
            case 't':
                mode[TEST] = 1;
                mode[IPC] = 0;
                break;
            default:
                perror("Unknown option");
                return EXIT_FAILURE;
        }
    }

    for (int i = optind; i < argc; i++) {
        if (strlen(filter) + strlen(argv[i]) + 2 <= sizeof(filter)) {
            strcat(filter, argv[i]);
            strcat(filter, " ");
        }
    }

    handle = CreatePcapHandle(device, filter);
    if (handle == NULL) {
        return EXIT_FAILURE;
    }

    linkhdrlen = GetLinkHeaderLen(handle);
    if (linkhdrlen == 0) {
        pcap_close(handle);
        return EXIT_FAILURE;
    }

    if (mode[VERBOSE]) {
        signal(SIGINT, StopCapture);
        signal(SIGTERM, StopCapture);
        signal(SIGQUIT, StopCapture);
        if (pcap_loop(handle, count, packet_handler, NULL) < 0) {
            perror("pcap_loop()");
            pcap_close(handle);
            return EXIT_FAILURE;
        }
        StopCapture();
    } else if (mode[IPC]) {
        signal(SIGINT, StopCaptureIPC);
        signal(SIGTERM, StopCaptureIPC);
        signal(SIGQUIT, StopCaptureIPC);
        if (pcap_loop(handle, count, packet_handler_IPC, NULL) < 0) {
            perror("pcap_loop()");
            pcap_close(handle);
            return EXIT_FAILURE;
        }
        StopCaptureIPC();
    } else {
        if (pcap_loop(handle, count, packet_handler_TEST, NULL) < 0) {
            perror("pcap_loop()");
            pcap_close(handle);
            return EXIT_FAILURE;
        }
        // 0 if test good, 255 if test bad
        return StopCaptureTest();
    }

    return EXIT_SUCCESS;
}
