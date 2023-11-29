package main

import (
	"time"

	"github.com/md-aryan-patel/projectx/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trRemote.Connect(trLocal)
	trLocal.Connect(trRemote)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello world!"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts {
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}