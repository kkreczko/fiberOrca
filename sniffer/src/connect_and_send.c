#include "connect_and_send.h"
#include "common.h"

#include <sys/socket.h>
#include <sys/un.h>

int CreateSocket() {
    const int sock = socket(AF_UNIX, SOCK_STREAM, 0);
    if (sock < 0) {
        perror("Socket creation");
        return -1;
    }

    struct sockaddr_un addr = {0};
    addr.sun_family = AF_UNIX;
    strcpy(addr.sun_path, SOCKET_PATH);

    if (connect(sock, (struct sockaddr *) &addr, sizeof(addr)) < 0) {
        perror("Socket connection\n");
        close(sock);
        return -1;
    }

    return sock;
}

int SendData(const int sock, const char* data, const size_t len) {
    const ssize_t sent = send(sock, data, len, 0);
    if (sent < 0) {
        perror("Socket data send");
        return -1;
    }

    return 0;
}
