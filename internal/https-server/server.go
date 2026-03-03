package httpsserver

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type HTTPSServer struct {
	config *Config
}

func NewHTTPSServer(config *Config) *HTTPSServer {
	return &HTTPSServer{
		config: config,
	}
}

func (s *HTTPSServer) handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	logRequest(r, 200)
	http.ServeFile(w, r, "internal/https-server/dashboard.html")
}

func (s *HTTPSServer) metricsAPIHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, 200)
	
	metrics := map[string]interface{}{
		"status": "online",
		"uptime": "0h 0m",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func (s *HTTPSServer) Start() {
	http.HandleFunc("/", s.handler)
	http.HandleFunc("/api/metrics", s.metricsAPIHandler)
	
	fs := http.FileServer(http.Dir("cache"))
	http.Handle("/cache/", http.StripPrefix("/cache/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r, 200)
		fs.ServeHTTP(w, r)
	})))

	go func() {
		httpPort := fmt.Sprintf(":%d", s.config.HTTPS.HTTPPort)
		color.Green("[HTTP] Listening on port %d...", s.config.HTTPS.HTTPPort)
		if err := http.ListenAndServe(httpPort, nil); err != nil {
			color.Red("[HTTP] Error: %v", err)
		}
	}()

	go func() {
		httpsPort := fmt.Sprintf(":%d", s.config.HTTPS.Port)
		color.Green("[HTTPS] Listening on port %d...", s.config.HTTPS.Port)
		
		certFile := s.config.HTTPS.CertFile
		keyFile := s.config.HTTPS.KeyFile
		
		if certFile == "" || keyFile == "" {
			color.Yellow("[HTTPS] No cert/key specified, generating self-signed certificate...")
			certFile, keyFile = s.generateSelfSignedCert()
		}
		
		server := &http.Server{
			Addr: httpsPort,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		}
		
		if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
			color.Red("[HTTPS] Error: %v", err)
		}
	}()
	
	color.Cyan("[SERVER] HTTPS server started successfully")
	select {}
}

func logRequest(r *http.Request, statusCode int) {
	method := r.Method
	path := r.URL.Path
	
	var statusColor func(format string, a ...interface{}) string
	switch {
	case statusCode >= 200 && statusCode < 300:
		statusColor = color.GreenString
	case statusCode >= 300 && statusCode < 400:
		statusColor = color.CyanString
	case statusCode >= 400 && statusCode < 500:
		statusColor = color.YellowString
	default:
		statusColor = color.RedString
	}
	
	log.Printf("[%s] %s %s - %s",
		r.RemoteAddr,
		method,
		path,
		statusColor("%d", statusCode),
	)
}
