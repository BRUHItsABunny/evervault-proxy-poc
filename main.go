package main

import (
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	gokhttp_client "github.com/BRUHItsABunny/gOkHttp/client"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/fetch", fetch)
	http.ListenAndServe(":80", nil)
}

func fetch(w http.ResponseWriter, req *http.Request) {
	// get parameters http headers x-proxy and x-fetch
	proxyURL := req.Header.Get("x-proxy") // url format: proto://user:pass@host:port
	fetchURL := req.Header.Get("x-fetch") // url
	if fetchURL == "" {
		// Default httpbin
		fetchURL = "https://httpbin.proxyman.app/get"
	}

	opts := []gokhttp_client.Option{}
	if proxyURL != "" {
		// if proxy is populated, try to use it
		opts = append(opts, gokhttp_client.NewProxyOption(proxyURL))
	}

	// Create client with options, in this case proxy option
	hClient, err := gokhttp.NewHTTPClient(opts...)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// Make a get request from client
	resp, err := hClient.Get(fetchURL)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// Copy response to response writer and close body.
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
	return
}
