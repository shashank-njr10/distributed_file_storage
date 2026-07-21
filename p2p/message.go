package p2p
import "net"

//Message hold any arbitrary data that is being
//sent over each transport between two nodes
//in the netwrork. The data is encoded and decoded using the
//transport specific encoder and decoder.
type RPC struct {
	From net.Addr
	Payload []byte 
}