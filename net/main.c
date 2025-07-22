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
  FILE *logFile = NULL;

  filter_t userFilter = {};
  struct sockaddr socketAddr;

  int sockFd, socketAddrLen, buffLen;
  uint8_t *buffer = (uint8_t *)malloc(65536);
  memset(buffer, 0, 65536);

  sockFd = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ALL));

  if (sockFd == -1)
    EXIT_MSG("NET: Failed to create raw socket");

  while (1) {
    static struct option long_options[] = {
      {"Source IP address", required_argument, NULL, 's'},
      {"Destination IP address", required_argument, NULL, 'd'},
      {"Source port", required_argument, NULL, 'p'},
      {"Destination port", required_argument, NULL, 't'},
      {"Source network interface", required_argument, NULL, 'i'},
      {"Destination network interface", required_argument, NULL, 'g'},
      {"Log file location", required_argument, NULL, 'o'},
      {"Catch TCP packets", no_argument, NULL, 'm'},
      {"Catch UDP packets", no_argument, NULL, 'n'},
      {0, 0, 0, 0}
    };

    int opt = getopt_long(argc, argv, "mns::d:p:t:i:g:o", long_options, NULL);

    if (opt == -1)
      break;

    switch (c) {

    }
  }
}
