package main

import (
	. "org.jboss.pnc.domain-proxy/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	domainProxyServer := NewDomainProxyServer()
	ready := make(chan bool)
	domainProxyServer.Start(ready)
	<-ready
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	domainProxyServer.Stop()
}
