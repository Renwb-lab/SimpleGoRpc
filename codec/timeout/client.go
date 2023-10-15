package timeout

import (
	clientPkg "SimpleGoRpc/codec/client"
	"SimpleGoRpc/codec/service"
)

// Dial connects to an RPC server at the specified network address
func Dial(network, address string, opts ...*service.Option) (*clientPkg.Client, error) {
	return clientPkg.DialTimeout(clientPkg.NewClient, network, address, opts...)
}
