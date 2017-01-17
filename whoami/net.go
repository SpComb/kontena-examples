package main

import (
  "fmt"
  "net"
  "net/http"
)

type NetworkInfo struct {
  LocalAddr net.Addr
  RemoteAddr net.Addr
}

func (networkInfo *NetworkInfo) fromConn(netConn net.Conn) error {
  networkInfo.LocalAddr = netConn.LocalAddr()
  networkInfo.RemoteAddr = netConn.RemoteAddr()

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
