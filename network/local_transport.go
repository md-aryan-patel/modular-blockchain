package network

import (
	"fmt"
	"sync"
)

type LocalTransprot struct {
	addr NetAddr
	consumeCh chan RPC
	lock sync.RWMutex
	peers map[NetAddr]*LocalTransprot
}

func NewLocalTransport(_addr NetAddr) Transport {
	return &LocalTransprot{
		addr: _addr,
		consumeCh: make(chan RPC, 1024),
		peers: make(map[NetAddr]*LocalTransprot),
	}
}

func (t* LocalTransprot) Consume() <- chan RPC {
	return t.consumeCh
}

func(t* LocalTransprot) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	
	t.peers[tr.Addr()] = tr.(*LocalTransprot)
	return nil 
}

func(t* LocalTransprot) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	peer.consumeCh <- RPC{
		From: t.addr,
		Payload: payload,
	}

	return nil
}

func(t* LocalTransprot) Addr() NetAddr {
	return t.addr
}