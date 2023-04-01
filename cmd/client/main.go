package main

import (
	"fmt"
	"github.com/kirillgrachoff/go-quic-potato/pkg/quicclient"
	"io"
	"log"
	"os"
)

func main() {
	client := quicclient.NewQuicClient("https://localhost:8080/catalog", false)
	resp, err := client.Get()
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Printf("status: %s\n", resp.Status)
	io.Copy(os.Stdout, resp.Body)
}
