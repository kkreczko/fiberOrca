#ifndef INCLUDE_CONNECT_AND_SEND_H
#define INCLUDE_CONNECT_AND_SEND_H

#include <pcap.h>

#define SOCKET_PATH "/tmp/fiber_orcas.sock"
#define BUFFER_SIZE 4096

int create_socket();
int send_data(int sock, const char* data, size_t len);

#endif // INCLUDE_CONNECT_AND_SEND_H