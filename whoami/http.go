package main

import (
  "encoding/json"
  "fmt"
  "net"
  "net/http"
  "log"
)

type HTTPHeaders map[string]string

func makeHttpHeaders(netHttpHeaders http.Header) HTTPHeaders {
  var httpHeaders = make(HTTPHeaders)

  for header, values := range netHttpHeaders {
    httpHeaders[header] = values[0]
  }

  return httpHeaders
}

type HTTPInfo struct {
  Headers HTTPHeaders
  Host string
}

func (httpInfo *HTTPInfo) fromHTTP(r *http.Request) error {
  httpInfo.Headers = makeHttpHeaders(r.Header)
  httpInfo.Host = r.Host

  return nil
}

func (networkInfo *NetworkInfo) fromHttp(httpRequest *http.Request) error {
  if httpRequest.RemoteAddr == "" {

  } else if tcpAddr, err := net.ResolveTCPAddr("tcp", httpRequest.RemoteAddr); err != nil {
    return fmt.Errorf("net.ResolveTCPAddr %v: %v", httpRequest.RemoteAddr, err)
  } else {
    networkInfo.RemoteAddr = tcpAddr
  }

  // this address is just the server listen address :(
  // https://github.com/golang/go/issues/6732#issuecomment-273164492
  if localAddr, ok := httpRequest.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
    networkInfo.LocalAddr = localAddr
  }

  return nil
}

type Whoami struct {
	OS      OSInfo
  Network NetworkInfo
  HTTP    HTTPInfo
}


func (whoami *Whoami) fromHTTP(r *http.Request) error {
  if err := whoami.OS.fromOS(); err != nil {
    return err
  }
  if err := whoami.Network.fromHttp(r); err != nil {
    return err
  }
  if err := whoami.HTTP.fromHTTP(r); err != nil {
    return err
  }

  return nil
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
  var whoami Whoami

  if err := whoami.fromHTTP(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)

  if err := json.NewEncoder(w).Encode(whoami); err != nil {
    log.Printf("json.Encode: %v", err)
	}
}

func httpMain(listen string) {
  log.Printf("http listen %v", listen)

  http.HandleFunc("/", httpHandler)

  if err := http.ListenAndServe(listen, nil); err != nil {
    log.Fatalf("http.ListenAndServe %v: %v", err)
  } else {
    log.Printf("http listen done")
  }
}
