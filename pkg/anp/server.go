package anp

import (
	"context"
	"net"
	"sync"
)

type Server struct {
	l    net.Listener
	p    sync.Pool
	sess map[net.Addr]*Client
	mu   sync.RWMutex

	cb Callbacks
}

type Callbacks interface {
	OnConnect(*Client)
	OnClose(*Client)
	OnError(*Client, error)
}

func NewServer() *Server {
	return &Server{
		sess: make(map[net.Addr]*Client),
		p: sync.Pool{
			New: func() any {
				return NewClient()
			},
		},
	}
}

func (s *Server) Listen(ctx context.Context, network, addr string) error {
	lc := net.ListenConfig{}
	l, err := lc.Listen(ctx, network, addr)
	if err != nil {
		return err
	}
	s.l = l
	return nil
}

func (s *Server) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := s.l.Accept()
			if err != nil {
				return err
			}
			go s.clientHandle(ctx, conn)
		}
	}
}

func (s *Server) clientHandle(ctx context.Context, conn net.Conn) {
	// client session setup
	c := s.p.Get().(*Client)
	defer s.p.Put(c)
	c.Init(ctx, conn)
	// map by address
	s.clientAdd(c)
	defer s.clientDel(c)
	// callbacks
	if s.cb != nil {
		s.cb.OnConnect(c)
		defer s.cb.OnClose(c)
	}
	// send io
	go c.RunSend()
	// recv io
	if err := c.RunRecv(); err != nil {
		if s.cb != nil {
			s.cb.OnError(c, err)
		}
	}
}

func (s *Server) clientAdd(c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sess[c.conn.RemoteAddr()] = c
}

func (s *Server) clientDel(c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sess, c.conn.RemoteAddr())
}
