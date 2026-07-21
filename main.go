package main

import (
	"log"
	"github.com/shashank-njr10/distributed_file_storage/p2p"
	"fmt"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts {
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		OnPeer: OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()


	if err:= tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}