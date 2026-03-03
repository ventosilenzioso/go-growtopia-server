package main

import (
	"fmt"
	"log"
	
	"project/internal/enet"
	httpsserver "project/internal/https-server"
)

func main() {
	fmt.Println("  Starting Server")	
	config, err := httpsserver.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println("\n[CONFIG] Configuration loaded successfully")
	fmt.Printf("[CONFIG] ENet Port: %d\n", config.ENet.Port)
	fmt.Printf("[CONFIG] ENet Max Peers: %d\n", config.ENet.MaxPeers)
	fmt.Printf("[CONFIG] HTTPS Port: %d\n", config.HTTPS.Port)
	fmt.Printf("[CONFIG] HTTP Port: %d\n", config.HTTPS.HTTPPort)
	
	fmt.Println("\n[MAIN] Starting HTTPS server in background...")
	go func() {
		httpsSrv := httpsserver.NewHTTPSServer(config)
		httpsSrv.Start()
	}()
	fmt.Println("[MAIN] HTTPS server started in goroutine")
	fmt.Println("\n[MAIN] Initializing ENet server...")
	enetSrv := enet.CreateServer("0.0.0.0", config.ENet.Port, config.ENet.MaxPeers)
	defer enetSrv.Close()
	fmt.Println("  All servers started successfully!")
	fmt.Printf("  HTTP:  http://localhost:%d\n", config.HTTPS.HTTPPort)
	fmt.Printf("  HTTPS: https://localhost:%d\n", config.HTTPS.Port)
	fmt.Printf("  ENet:  UDP port %d\n", config.ENet.Port)	
	fmt.Println("[MAIN] Starting ENet event loop (this will block)...")
	enetSrv.Listen()
}
