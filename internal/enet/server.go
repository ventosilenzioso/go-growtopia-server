package enet

/*
#cgo CFLAGS: -I../../include
#include "enet/enet.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

type Server struct {
	host *C.ENetHost
}

func CreateServer(ip string, port int, maxPeers int) *Server {
	fmt.Println("[ENET] Initializing ENet...")
	if C.enet_initialize() != 0 {
		panic("Failed to initialize ENet")
	}
	fmt.Println("[ENET] ENet initialized successfully")

	var addr C.ENetAddress
	addr.port = C.enet_uint16(port)
	if ip == "" || ip == "0.0.0.0" {
		addr.host = C.ENET_HOST_ANY
		fmt.Printf("[ENET] Binding to 0.0.0.0:%d (all interfaces)\n", port)
	} else {
		cip := C.CString(ip)
		defer C.free(unsafe.Pointer(cip))
		C.enet_address_set_host(&addr, cip)
		fmt.Printf("[ENET] Binding to %s:%d\n", ip, port)
	}

	fmt.Printf("[ENET] Creating host with max %d peers...\n", maxPeers)
	host := C.enet_host_create(&addr, C.size_t(maxPeers), 2, 0, 0)
	if host == nil {
		panic("Failed to create ENet server")
	}
	fmt.Println("[ENET] Host created successfully")

	host.checksum = C.ENetChecksumCallback(C.enet_crc32)
	host.usingNewPacketForServer = 1
	fmt.Println("[ENET] Checksum and packet settings configured")
	fmt.Println("[ENET] Server is ready and listening for connections")

	return &Server{host: host}
}

func (s *Server) Listen() {
	var event C.ENetEvent
	fmt.Println("[ENET] Starting event loop...")
	fmt.Println("[ENET] Waiting for connections...")
	
	for {
		for C.enet_host_service(s.host, &event, 1000) > 0 {
			peer := event.peer
			
			switch event._type {
			case C.ENET_EVENT_TYPE_CONNECT:
				var clientIP [16]C.char
				C.enet_address_get_host_ip(&peer.address, &clientIP[0], 16)
				ipStr := C.GoString(&clientIP[0])
				peerID := int(peer.connectID)
				
				fmt.Printf("[ENET] ========================================\n")
				fmt.Printf("[ENET] NEW CONNECTION\n")
				fmt.Printf("[ENET] IP: %s\n", ipStr)
				fmt.Printf("[ENET] Port: %d\n", peer.address.port)
				fmt.Printf("[ENET] Peer ID: %d\n", peerID)
				fmt.Printf("[ENET] ========================================\n")
				
			case C.ENET_EVENT_TYPE_RECEIVE:
				fmt.Printf("[ENET] Received packet from peer (size: %d bytes)\n", event.packet.dataLength)
				C.enet_packet_destroy(event.packet)
				
			case C.ENET_EVENT_TYPE_DISCONNECT:
				fmt.Printf("[ENET] Client disconnected (data: %d)\n", event.data)
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (s *Server) Close() {
	if s.host != nil {
		C.enet_host_destroy(s.host)
		C.enet_deinitialize()
		fmt.Println("[ENET] Server closed")
	}
}
