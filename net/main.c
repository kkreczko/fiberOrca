#include <asm-generic/socket.h>
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

#include "include/dbg_helpers.h"

#define BUFFSIZ 2048

// BELOW HAS TO GO VERY STUPID
#define IF_NAME "wlp4s0"

char *transport_protocol(unsigned int code)
{
  switch(code)
  {
    case 1: return "icmp";
    case 2: return "igmp";
    case 6: return "tcp";
    case 17: return "udp";
    default: return "unknown";
  }
}

int main(int argc, char **argv)
{
  int sock, n;
  char buffer[BUFFSIZ];
  unsigned char *iphead, *ethhead;

  if ((sock = socket(PF_PACKET, SOCK_RAW, htons(ETH_P_IP))) == -1)
    EXIT_MSG("NET: failed to create socket")

  const char *opt;

  // TODO IMPORTANT
  // it has to be found automatically rather than hardcoded
  opt = IF_NAME;
  
  if (setsockopt(sock, SOL_SOCKET, SO_BINDTODEVICE, opt, strlen(opt) + 1) == -1) 
  {
    close(sock);
    EXIT_MSG("NET: failed to bind to device")
  }

  struct ifreq ethreq;

  strncpy(ethreq.ifr_name, IF_NAME, IF_NAMESIZE);
  if (ioctl(sock, SIOCGIFFLAGS, &ethreq) == -1) 
  {
    close(sock);
    EXIT_MSG("NET: failed to set device name");
  }

  ethreq.ifr_flags |= IFF_PROMISC;
  if (ioctl(sock, SIOCGIFFLAGS, &ethreq) == -1) 
  {
    close(sock);
    EXIT_MSG("NET: failed to set device into promiscous mode");
  }

  close(sock);
  return 0;
}
