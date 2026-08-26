package main

import "github.com/shashank-njr10/distributed_file_storage/p2p"
import "log"
// import "time"



func makeServer(listenAddr string, nodes ...string) *FileServer {
		tcpTransportOpts := p2p.TCPTransportOpts {
		ListenAddr: listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		//TODO on peer func
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts {
		StorageRoot: listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport: tcpTransport,
		BootStrapNodes:  nodes,
	}

	return NewFileServer(fileServerOpts)
}

func main() {
	s1 := makeServer(":3000","")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()

	s2.Start()

}