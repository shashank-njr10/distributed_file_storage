package p2p

//handshakeFunc is 
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }