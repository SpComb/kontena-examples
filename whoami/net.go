package main

import (
  "net"
)

type NetworkInfo struct {
  LocalAddr net.Addr
  RemoteAddr net.Addr
}

func (networkInfo *NetworkInfo) fromNet(netConn net.Conn) error {
  networkInfo.LocalAddr = netConn.LocalAddr()
  networkInfo.RemoteAddr = netConn.RemoteAddr()

  return nil
}
