package main

import (
	. "org.jboss.pnc.domain-proxy/pkg/client"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	domainProxyClient := NewDomainProxyClient()
	ready := make(chan bool)
	domainProxyClient.Start(ready)
	<-ready
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	domainProxyClient.Stop()
}
