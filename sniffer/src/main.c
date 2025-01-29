#include "capture.h"
#include "packet_handler.h"
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

//TODO switch from printing to ICP communication with bubbletea process
int main(int argc, char* argv[]) {
    char device[256] = {0};
    char filter[256] = {0};
    int count = 0;
    int is_verbose = 0;
    int opt;

    while ((opt = getopt(argc, argv, "hi:n:")) != -1) {
        switch (opt) {
            case 'h':
                printf("Usage: %s [-i interface] [-n count] [filter]\nexample: -n 64 tcp port 80", argv[0]);
                return EXIT_SUCCESS;
            case 'i':
                strncpy(device, optarg, sizeof(device) - 1);
                break;
            case 'n':
                count = atoi(optarg);
                break;
            case 'v':
                is_verbose = 1;
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

    if (is_verbose) {
        if (pcap_loop(handle, count, packet_handler, NULL) < 0) {
            fprintf(stderr, "pcap_loop failed: %s\n", pcap_geterr(handle));
            pcap_close(handle);
            return EXIT_FAILURE;
        }
        stop_capture();
    } else {
        if (pcap_loop(handle, count, packet_handler_ICP, NULL) < 0) {
            fprintf(stderr, "pcap_loop failed: %s\n", pcap_geterr(handle));
            pcap_close(handle);
            return EXIT_FAILURE;
        }
        stop_capture_ICP();
    }

    return EXIT_SUCCESS;
}