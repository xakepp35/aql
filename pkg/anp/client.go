package anp

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/xakepp35/aql/pkg/aql"
	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf"
)

type Client struct {
	conn  net.Conn
	netvm *aql.VM
	inbox chan any
}

func NewClient() *Client {
	return &Client{
		netvm: aql.New(),
		inbox: make(chan any, 16),
	}
}

func Dial(ctx context.Context, network, addr string) (*Client, error) {
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}
	c := NewClient()
	c.Init(ctx, conn)
	go c.RunSend()
	go c.RunRecv()
	return c, nil
}

func (s *Client) Init(ctx context.Context, conn net.Conn) {
	s.netvm.SetCtx(context.WithCancel(ctx))
	s.conn = conn
}

func (s *Client) RunRecv() error {
	defer s.conn.Close()
	for {
		select {
		case <-s.netvm.Context.Done():
			return s.netvm.Ctx().Err()

		// receive
		default:
			if err := s.Recv(); err != nil {
				return err
			}
			s.netvm.Run()
			if s.netvm.Err != nil {
				fmt.Println("netvm: " + s.netvm.Err.Error())
			}
		}
	}
}

func (s *Client) RunSend() error {
	for {
		select {
		case <-s.netvm.Context.Done():
			return s.netvm.Ctx().Err()

		// send
		case msg, ok := <-s.netvm.SendStream:
			if !ok {
				return errors.New("send stream closed")
			}
			code, err := json.Marshal(msg)
			if err != nil {
				return err
			}
			b := make([]byte, 5+len(code))
			b[0] = byte(op.Halt)
			binary.LittleEndian.PutUint32(b[1:], uint32(len(code)))
			copy(b[5:], code)
			if err := s.Send(b); err != nil {
				return err
			}
		}
	}
}

func (s *Client) Recv() error {
	var sizeBuf [4]byte
	if _, err := io.ReadFull(s.conn, sizeBuf[:]); err != nil {
		return err
	}
	size := binary.BigEndian.Uint32(sizeBuf[:])
	if len(s.netvm.Emit) < int(size) {
		s.netvm.Emit = make(asf.Emitter, size)
	}
	s.netvm.Emit = s.netvm.Emit[:size]
	if _, err := io.ReadFull(s.conn, s.netvm.Emit); err != nil {
		return err
	}
	return nil
}

func (s *Client) Send(code []byte) error {
	var sizeBuf [4]byte
	binary.BigEndian.PutUint32(sizeBuf[:], uint32(len(code)))
	if err := s.send_loop(sizeBuf[:]); err != nil {
		return err
	}
	if err := s.send_loop(code); err != nil {
		return err
	}
	return nil
}

func (s *Client) send_loop(b []byte) error {
	var n int
	for n < len(b) {
		nn, err := s.conn.Write(b[n:])
		if err != nil {
			return err
		}
		n += nn
	}
	return nil
}
