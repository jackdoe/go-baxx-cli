package util

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

func ReadKey(from string) ([]byte, error) {
	h := sha256.New()
	f, err := os.Open(from)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}

	shakey := h.Sum(nil)
	return shakey, nil
}

func Exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}

func ExitHttp(resp *http.Response) {
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		Exit(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", string(b))
	os.Exit(1)
}
