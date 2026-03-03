package httpsserver

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

func (s *HTTPSServer) generateSelfSignedCert() (string, string) {
	certDir := "internal/https-server/certs"
	os.MkdirAll(certDir, 0755)

	certFile := filepath.Join(certDir, "certificate.pem")
	keyFile := filepath.Join(certDir, "privatekey.pem")

	if _, err := os.Stat(certFile); err == nil {
		return certFile, keyFile
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	notBefore := time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)
	notAfter := notBefore.Add(10 * 365 * 24 * time.Hour)

	serialNumber, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country:            []string{"ID"},
			Province:           []string{"Makassar"},
			Locality:           []string{"QWERTY"},
			Organization:       []string{"QWerty Ltd"},
			OrganizationalUnit: []string{""},
			CommonName:         "www.growtopia1.com",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)

	certOut, _ := os.Create(certFile)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create(keyFile)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()

	return certFile, keyFile
}
