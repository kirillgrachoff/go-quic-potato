/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/kirillgrachoff/go-quic-potato/pkg/quicclient"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// runClientCmd represents the runClient command
var runClientCmd = &cobra.Command{
	Use:   "runClient",
	Short: "A brief description of your command",
	Long:  `./go-quic-potato runClient https://localhost:8080/catalog -i`,
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		client := quicclient.NewQuicClient(url, insecure)
		resp, err := client.Get()
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Printf("status: %s\n", resp.Status)
		io.Copy(os.Stdout, resp.Body)
	},
}

var insecure bool

func init() {
	rootCmd.AddCommand(runClientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runClientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runClientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runClientCmd.Flags().BoolVarP(&insecure, "insecure", "i", false, "not to check certificates")
}
