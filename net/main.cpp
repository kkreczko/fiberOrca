#include <cstdint>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <stdint.h>
#include <getopt.h>
#include <errno.h>
#include <sys/ioctl.h>
#include <linux/if_packet.h>
#include <net/if.h>
#include <netinet/if_ether.h>
#include <netinet/ip.h>
#include <netinet/udp.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/tcp.h>

#define EXIT_MSG(msg) do { perror(msg); exit(EXIT_FAILURE); } while(0)

typedef struct {
  uint8_t tProto;
  char *srcIp;
  char *dstIp;
  uint16_t srcPort;
  uint16_t dstPort;
  char *srcIfName;
  char *dstIfName;
  uint8_t srcMac[6];
  uint8_t dstMac[6];
} filter_t;

struct sockaddr_in sourceAddr, destAddr;

int main(int argc, char **argv)
{
  int c;
  char log[255];
  FILE *logfile = NULL;

  filter_t userFilter = {};
}
