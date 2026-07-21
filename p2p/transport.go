package p2p

// Peer is an interface that represents the remote nodes in the network.
type Peer interface {

}
// Transport is anything that handles the communication between nodes in the network.
// this can be TCP, UDP or webscokets. The transport layer is responsible for sending and receiving messages between nodes.
type Transport interface {
	ListenAndAccept() error
}
