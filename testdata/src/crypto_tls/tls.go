package crypto

import (
	"context"
	"crypto/tls"
	"net"
)

func _() {
	ctx := context.Background()

	netDialer := &net.Dialer{}

	tlsConfig := &tls.Config{}

	// dialers
	tls.Dial("tcp", "localhost:8080", tlsConfig)                      // want `crypto/tls\.Dial must not be called. use \(\*crypto/tls\.Dialer\)\.DialContext`
	tls.DialWithDialer(netDialer, "tcp", "localhost:8080", tlsConfig) // want `crypto/tls\.DialWithDialer must not be called. use \(\*crypto/tls\.Dialer\)\.DialContext with NetDialer`

	tlsDialer := &tls.Dialer{
		Config:    tlsConfig,
		NetDialer: netDialer,
	}
	tlsDialer.DialContext(ctx, "tcp", "localhost:8080")

	// connection
	tlsConn := &tls.Conn{}
	tlsConn.Handshake() // want `\(\*crypto/tls\.Conn\)\.Handshake must not be called. use \(\*crypto/tls\.Conn\).HandshakeContext`

	tlsConn.HandshakeContext(ctx)
}
