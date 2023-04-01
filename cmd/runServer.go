package cmd

import (
	"github.com/kirillgrachoff/go-quic-potato/pkg/quicserver"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

// runServerCmd represents the runServer command
var runServerCmd = &cobra.Command{
	Use:   "runServer",
	Short: "Run a QUIC server",
	Long:  `./go-quic-potato runServer -f ./testdata/catalog.json`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath := cmd.Flag("file").Value.String()
		addr := cmd.Flag("addr").Value.String()

		s := quicserver.QuicServer{
			FilePath: filePath,
		}
		f, err := os.Open(s.FilePath)
		if err != nil {
			log.Fatalln(err)
		}
		buf, _ := io.ReadAll(f)
		log.Println(string(buf))
		f.Close()
		log.Fatalln(s.ListenAndServe(addr))
	},
}

func init() {
	rootCmd.AddCommand(runServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runServerCmd.Flags().StringP("file", "f", "./file.txt", "file to return")
	runServerCmd.Flags().StringP("addr", "a", "localhost:8080", "address to serve")
}
