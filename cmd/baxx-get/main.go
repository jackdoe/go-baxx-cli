package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jackdoe/go-baxx-cli/util"
	"github.com/minio/sio"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "Example:\n  %s -k pass-file https://baxx.dev/io/$BAXX_TOKEN/file > file.downloaded\n\n", os.Args[0])
	}

	var pkeyFrom = flag.String("k", "", "file containing the encryption key, key is sha256(content of the file)")
	// add timeout support
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	client := &http.Client{}
	shakey, err := util.ReadKey(*pkeyFrom)
	if err != nil {
		util.Exit(err)
	}

	req, err := http.NewRequest("GET", flag.Arg(0), nil)
	resp, err := client.Do(req)
	if err != nil {
		util.Exit(err)
	}
	if resp.StatusCode != 200 {
		util.ExitHttp(resp)
	}
	reader, err := sio.DecryptReader(resp.Body, sio.Config{Key: shakey})
	if err != nil {
		util.Exit(err)
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		util.Exit(err)
	}
}
