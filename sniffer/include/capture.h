#ifndef INCLUDE_CAPTURE_H
#define INCLUDE_CAPTURE_H

#include "common.h"

inline pcap_t *CreatePcapHandle(char *device, char *filter);
int GetLinkHeaderLen(pcap_t *handle);
void StopCapture();
void StopCaptureIPC();
int StopCaptureTest();

#endif // INCLUDE_CAPTURE_H
