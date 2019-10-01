package socketcan

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.einride.tech/can"
	"golang.org/x/sync/errgroup"
)

func TestDial_TCP(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	var g errgroup.Group
	g.Go(func() error {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		return conn.Close()
	})
	conn, err := Dial("tcp", lis.Addr().String())
	require.NoError(t, err)
	require.NoError(t, conn.Close())
	require.NoError(t, g.Wait())
}

func TestDialContext_TCP(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	var g errgroup.Group
	g.Go(func() error {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		return conn.Close()
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second)
	defer done()
	conn, err := DialContext(ctx, "tcp", lis.Addr().String())
	require.NoError(t, err)
	require.NoError(t, conn.Close())
	require.NoError(t, g.Wait())
}

func TestConn_TransmitReceiveTCP(t *testing.T) {
	// Given: A TCP listener that writes a frame on an accepted connection
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	var g errgroup.Group
	frame := can.Frame{ID: 42, Length: 5, Data: can.Data{'H', 'e', 'l', 'l', 'o'}}
	g.Go(func() error {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		tr := NewTransmitter(conn)
		ctx, done := context.WithTimeout(context.Background(), time.Second)
		defer done()
		if err := tr.TransmitFrame(ctx, frame); err != nil {
			return err
		}
		return conn.Close()
	})
	// When: We connect to the listener
	ctx, done := context.WithTimeout(context.Background(), time.Second)
	defer done()
	conn, err := DialContext(ctx, "tcp", lis.Addr().String())
	require.NoError(t, err)
	rec := NewReceiver(conn)
	require.True(t, rec.Receive())
	require.False(t, rec.HasErrorFrame())
	require.Equal(t, frame, rec.Frame())
	require.NoError(t, conn.Close())
	require.NoError(t, g.Wait())
}
