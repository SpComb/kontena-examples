package main

import (
  "encoding/json"
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

type Whoami struct {
	OS      OSInfo
  Network NetworkInfo
  HTTP    HTTPInfo
}

func (whoami *Whoami) fromHTTP(r *http.Request) error {
  if err := whoami.OS.fromOS(); err != nil {
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
  w.Header().Set("Connection", "close")
  w.Header().Set("Transfer-Encoding", "identity")
  w.WriteHeader(200)

  // hijack to get Network.LocalAddr?
  if hijacker, ok := w.(http.Hijacker); !ok {
    if err := whoami.Network.fromHttp(r); err != nil {
      panic(err)
    }

    if err := json.NewEncoder(w).Encode(whoami); err != nil {
      panic(err)
  	}
  } else if conn, bufio, err := hijacker.Hijack(); err != nil {
    panic(err)
  } else {
    defer conn.Close()

    // get real network info from conn
    if err := whoami.Network.fromConn(conn); err != nil {
      panic(err)
    }

    // write JSON responses
    if err := json.NewEncoder(bufio).Encode(whoami); err != nil {
      panic(err)
    }

    if err := bufio.Flush(); err != nil {
      panic(err)
    }
  }
}

func httpMain(listen string, static string) {
  log.Printf("http listen %v", listen)

  http.Handle("/", http.FileServer(http.Dir(static)))
  http.HandleFunc("/api", httpHandler)

  if err := http.ListenAndServe(listen, nil); err != nil {
    log.Fatalf("http.ListenAndServe %v: %v", err)
  } else {
    log.Printf("http listen done")
  }
}
