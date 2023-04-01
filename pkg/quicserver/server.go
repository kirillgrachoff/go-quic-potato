package quicserver

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/gorilla/mux"
	"github.com/quic-go/quic-go/http3"
	"io"
	"log"
	"math/big"
	"net/http"
)

type QuicServer struct {
	FilePath string
}

//type ListenerRepaired struct {
//	quic.Listener
//}
//
//func (l *ListenerRepaired) Accept() (net.Conn, error) {
//	return l.Listener.Accept(context.Background())
//}

func (q *QuicServer) ListenAndServe(addr string) error {
	mux := mux.NewRouter()

	mux.Path("/catalog").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buf, _ := io.ReadAll(r.Body)
		log.Printf("new request: %s\n", string(buf))
		http.ServeFile(w, r, q.FilePath)
	})

	server := http3.Server{
		Addr:      addr,
		TLSConfig: generateTLSConfig(),
		Handler:   mux,
	}

	return server.ListenAndServe()
}

// Set up a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
}
