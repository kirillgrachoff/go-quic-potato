package quicclient

import (
	"crypto/tls"
	"github.com/quic-go/quic-go/http3"
	"net/http"
)

type QuicClient struct {
	client http.Client
	url    string
}

func NewQuicClient(url string, insecure bool) *QuicClient {
	return &QuicClient{
		client: http.Client{
			Transport: &http3.RoundTripper{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
		url: url,
	}
}

func (q *QuicClient) Get() (*http.Response, error) {
	return q.client.Get(q.url)
}
