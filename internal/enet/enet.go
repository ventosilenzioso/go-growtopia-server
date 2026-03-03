package enet

/*
#cgo CFLAGS: -I../../include
#cgo windows LDFLAGS: -lws2_32 -lwinmm
#include "enet/enet.h"
#include "../../include/enet/callbacks.c"
#include "../../include/enet/compress.c"
#include "../../include/enet/host.c"
#include "../../include/enet/list.c"
#include "../../include/enet/packet.c"
#include "../../include/enet/peer.c"
#include "../../include/enet/protocol.c"
#include "../../include/enet/win32.c"
*/
import "C"
