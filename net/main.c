#include <sys/types.h>
#include <sys/socket.h>
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
#include <net/if_arp.h>
#include <netinet/if_ether.h>
#include <netinet/ip.h>
#include <netinet/udp.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <linux/filter.h>

#include "include/dbg_helpers.h"

#define BUFFSIZ 2048

// BELOW HAS TO GO, VERY STUPID
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

  // TODO: We need to get it automatically from??? tcpdump??? :rofl:
  struct sock_filter BPF_code[] = {
    { 0x28, 0, 0, 0x0000000c },
    { 0x15, 0, 2, 0x00000800 },
    { 0x30, 0, 0, 0x00000017 },
    { 0x15, 6, 7, 0x00000006 },
    { 0x15, 0, 6, 0x000086dd },
    { 0x30, 0, 0, 0x00000014 },
    { 0x15, 3, 0, 0x00000006 },
    { 0x15, 0, 3, 0x0000002c },
    { 0x30, 0, 0, 0x00000036 },
    { 0x15, 0, 1, 0x00000006 },
    { 0x6, 0, 0, 0x00040000 },
    { 0x6, 0, 0, 0x00000000 },
  };

  struct sock_fprog Filter;
  Filter.len = sizeof(BPF_code)/sizeof(BPF_code[0]); 
  Filter.filter = BPF_code;

  if (setsockopt(sock, SOL_SOCKET, SO_ATTACH_FILTER, &Filter, sizeof(Filter)) < 0) {
        close(sock);
        EXIT_MSG("BPF: failed to set socket filter");
  } 

  while(1) {
    printf("-----------\n");
    n = recvfrom(sock, buffer, 2048, 0, NULL, NULL);
    printf("%d bytes read\n", n);

    if (n < 42) {
      perror("recvfrom():");
      printf("Incomplete packet (errno is %d)\n", errno);
      close(sock);
      exit(0);
    }

    ethhead = buffer;
    printf("Source MAC address: %.2x:%.2x:%.2x:%.2x:%.2x:%.2x\n",
      ethhead[0], ethhead[1], ethhead[2], ethhead[3], ethhead[4], ethhead[5]
    );
    printf("Destination MAC address: %.2x:%.2x:%.2x:%.2x:%.2x:%.2x\n",
      ethhead[6], ethhead[7], ethhead[8], ethhead[9], ethhead[10], ethhead[11]
    );

    iphead = buffer + 14; 

    if (*iphead==0x45) {
      printf("Source host %d.%d.%d.%d\n",
          iphead[12],iphead[13],
          iphead[14],iphead[15]);
      printf("Dest host %d.%d.%d.%d\n",
          iphead[16],iphead[17],
          iphead[18],iphead[19]);
      printf("Source,Dest ports %d,%d\n",
          (iphead[20]<<8)+iphead[21],
          (iphead[22]<<8)+iphead[23]);
      printf("Layer-4 protocol %s\n", transport_protocol(iphead[9]));
    }
  }
  
  close(sock);
  return 0;
}
