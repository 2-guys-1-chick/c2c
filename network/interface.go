package network

import (
	"github.com/2-guys-1-chick/c2c/network/packet"
	"net"
)

type Distributor interface {
	Distribute(packet *packet.Data) error
}

type PacketHandler interface {
	Handle(packet *packet.Data) error
}

type ByteDistributor interface {
	Distribute([]byte)
}

type DisconnectHandler interface {
	OnDisconnect(conn net.Conn)
}
