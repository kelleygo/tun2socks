package mirror

import (
	"github.com/kelleygo/tun2socks/v2/core/adapter"
	"github.com/kelleygo/tun2socks/v2/tunnel"
)

var _ adapter.TransportHandler = (*Tunnel)(nil)

type Tunnel struct{}

func (*Tunnel) HandleTCP(conn adapter.TCPConn) {
	tunnel.TCPIn() <- conn
}

func (*Tunnel) HandleUDP(conn adapter.UDPConn) {
	tunnel.UDPIn() <- conn
}
