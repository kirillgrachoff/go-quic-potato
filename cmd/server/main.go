package main

import (
	"io"
	"log"
	"os"

	"github.com/kirillgrachoff/go-quic-potato/pkg/quicserver"
)

func main() {
	s := quicserver.QuicServer{
		FilePath: "testdata/catalog.json",
	}
	f, err := os.Open(s.FilePath)
	if err != nil {
		log.Fatalln(err)
	}
	buf, _ := io.ReadAll(f)
	log.Println(string(buf))
	f.Close()
	log.Fatalln(s.ListenAndServe("localhost:8080"))
}
