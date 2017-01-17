package main

import (
  "fmt"
  "net"
  "net/http"
)

type AddrInfo struct {
  IP string
  Port int
  Name string
}

type NetworkInfo struct {
  LocalAddr AddrInfo
  RemoteAddr AddrInfo
}

func (addrInfo *AddrInfo) resolveDNS() error {
  if names, err := net.LookupAddr(addrInfo.IP); err != nil {
    return fmt.Errorf("net.LookupAddr %v: %v", addrInfo.IP, err)
  } else {
    for _, name := range names {
      addrInfo.Name = name
    }
  }

  return nil
}

func (addrInfo *AddrInfo) fromAddr(addr net.Addr) error {
  switch typeAddr := addr.(type) {
  case *net.TCPAddr:
    addrInfo.IP = typeAddr.IP.String()
    addrInfo.Port = typeAddr.Port
  default:
    return fmt.Errorf("Unknown addr type %T: %v", addr, addr)
  }

  return addrInfo.resolveDNS()
}

func (networkInfo *NetworkInfo) fromConn(netConn net.Conn) error {
  networkInfo.LocalAddr.fromAddr(netConn.LocalAddr())
  networkInfo.RemoteAddr.fromAddr(netConn.RemoteAddr())

  return nil
}

func (networkInfo *NetworkInfo) fromHttp(httpRequest *http.Request) error {
  if httpRequest.RemoteAddr == "" {

  } else if tcpAddr, err := net.ResolveTCPAddr("tcp", httpRequest.RemoteAddr); err != nil {
    return fmt.Errorf("net.ResolveTCPAddr %v: %v", httpRequest.RemoteAddr, err)
  } else {
    networkInfo.RemoteAddr.fromAddr(tcpAddr)
  }

  // this address is just the server listen address :(
  // https://github.com/golang/go/issues/6732#issuecomment-273164492
  if localAddr, ok := httpRequest.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
    networkInfo.LocalAddr.fromAddr(localAddr)
  } else {
    // not available
  }

  return nil
}
