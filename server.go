package main
import "github.com/shashank-njr10/distributed_file_storage/p2p"
import "io"
import "fmt"
import "log"


type FileServerOpts struct {
	StorageRoot string
	PathTransformFunc PathTransformFunc
	Transport 	p2p.Transport
	BootStrapNodes []string
}

type FileServer struct {
	FileServerOpts

	store *Store
	quitch chan struct{}
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storeOpts := StoreOpts {
		Root: opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store: NewStore(storeOpts),
		quitch: make(chan struct{}),
	}
}

func (s *FileServer) Store(key string, r io.Reader) error {
	return s.store.Write(key,r)
}

func (s *FileServer) Stop() {
	close(s.quitch)
}

func (s *FileServer) loop() {
	defer func() {
		log.Println("file server stopped due to user quit action")
		s.Transport.Close()
	}()
	for {
		select {
		case msg := <- s.Transport.Consume():
			fmt.Println(msg)
		case <- s.quitch:
			return 
		}
	}
}


func (s *FileServer) bootstrapNetwork() error {
	for _, addr :=  range s.BootStrapNodes {

		if len(addr) == 0 {
			continue
		}
		
		go func (addr string) {
			fmt.Println("attempting to connect with remote: ", addr)
			if err := s.Transport.Dial(addr); err != nil {
				log.Println("dial error: ", err)
			}
		} (addr)
	}

	return nil
}

func (s *FileServer) Start() error{
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}

	if len(s.BootStrapNodes) != 0 {
			s.bootstrapNetwork()
	}


	s.loop()

	return nil
}

