#ifndef INCLUDE_CONNECT_AND_SEND_H
#define INCLUDE_CONNECT_AND_SEND_H

#include <common.h>

#define SOCKET_PATH "/tmp/fiber_orcas.sock"
#define BUFFER_SIZE 1024

int CreateSocket();
int SendData(int sock, const char* data, size_t len);

#endif // INCLUDE_CONNECT_AND_SEND_H