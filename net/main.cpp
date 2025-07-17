#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/tcp.h>

int main()
{
  int raw_socket = socket(AF_INET, SOCK_RAW, IPPROTO_TCP);
  return 0;
}
