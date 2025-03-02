package quicconn

import (
	"context"
	"net"

	quic "github.com/quic-go/quic-go"
)

type server struct {
	quicServer quic.Listener
}

var _ net.Listener = &server{}

// Accept waits for and returns the next connection to the listener.
func (s *server) Accept() (net.Conn, error) {
	sess, err := s.quicServer.Accept(context.Background())
	if err != nil {
		return nil, err
	}
	qconn, err := newConn(sess)
	if err != nil {
		return nil, err
	}
	return qconn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (s *server) Close() error {
	return s.quicServer.Close()
}

// Addr returns the listener's network address.
func (s *server) Addr() net.Addr {
	return s.quicServer.Addr()
}
