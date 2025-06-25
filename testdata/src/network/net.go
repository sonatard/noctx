package network

import (
	"context"
	"net"
	"time"
)

var timeout = 10 * time.Second

func _() {
	ctx := context.Background()

	// listenConfig
	listenConfig := &net.ListenConfig{}

	net.Listen("tcp", "localhost:8080") // want `net\.Listen must not be called. use \(\*net\.ListenConfig\)\.Listen`
	listenConfig.Listen(ctx, "tcp", "localhost:8080")

	net.ListenPacket("udp", "localhost:8080") // want `net\.ListenPacket must not be called. use \(\*net\.ListenConfig\)\.ListenPacket`
	listenConfig.ListenPacket(ctx, "udp", "localhost:8080")

	// dialer
	net.Dial("tcp", "localhost:8080")                 // want `net\.Dial must not be called. use \(\*net\.Dialer\)\.DialContext`
	net.DialTimeout("tcp", "localhost:8080", timeout) // want `net\.DialTimeout must not be called. use \(\*net\.Dialer\)\.DialContext with \(\*net\.Dialer\)\.Timeout`

	dialer := &net.Dialer{}
	dialer.DialContext(ctx, "tcp", "localhost:8080")

	dialerTimeout := &net.Dialer{Timeout: timeout}
	dialerTimeout.DialContext(ctx, "tcp", "localhost:8080")

	// resolver
	net.LookupCNAME("example.com")              // want `net\.LookupCNAME must not be called. use \(\*net\.Resolver\)\.LookupCNAME with a context`
	net.LookupHost("example.com")               // want `net\.LookupHost must not be called. use \(\*net\.Resolver\)\.LookupHost with a context`
	net.LookupIP("example.com")                 // want `net\.LookupIP must not be called. use \(\*net\.Resolver\)\.LookupIPAddr with a context`
	net.LookupPort("tcp", "http")               // want `net\.LookupPort must not be called. use \(\*net\.Resolver\)\.LookupPort with a context`
	net.LookupSRV("http", "tcp", "example.com") // want `net\.LookupSRV must not be called. use \(\*net\.Resolver\)\.LookupSRV with a context`
	net.LookupMX("example.com")                 // want `net\.LookupMX must not be called. use \(\*net\.Resolver\)\.LookupMX with a context`
	net.LookupNS("example.com")                 // want `net\.LookupNS must not be called. use \(\*net\.Resolver\)\.LookupNS with a context`
	net.LookupTXT("example.com")                // want `net\.LookupTXT must not be called. use \(\*net\.Resolver\)\.LookupTXT with a context`
	net.LookupAddr("example.com")               // want `net\.LookupAddr must not be called. use \(\*net\.Resolver\)\.LookupAddr with a context`

	resolver := net.DefaultResolver
	resolver.LookupCNAME(ctx, "example.com")
	resolver.LookupHost(ctx, "example.com")
	resolver.LookupIPAddr(ctx, "example.com")
	resolver.LookupPort(ctx, "tcp", "http")
	resolver.LookupSRV(ctx, "http", "tcp", "example.com")
	resolver.LookupMX(ctx, "example.com")
	resolver.LookupNS(ctx, "example.com")
	resolver.LookupTXT(ctx, "example.com")
	resolver.LookupAddr(ctx, "example.com")
}
