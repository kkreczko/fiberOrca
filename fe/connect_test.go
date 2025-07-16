package main

import (
	"bytes"
	"github.com/charmbracelet/bubbles/list"
	"github.com/kruczys/fiberOrca/models"
	"io"
	"net"
	"testing"
	"time"
)

// mockConn is a minimal in-memory net.Conn for testing
type mockConn struct {
	io.Reader
}

func (m *mockConn) Read(b []byte) (int, error)         { return m.Reader.Read(b) }
func (m *mockConn) Write(b []byte) (int, error)        { return len(b), nil }
func (m *mockConn) Close() error                       { return nil }
func (m *mockConn) LocalAddr() net.Addr                { return nil }
func (m *mockConn) RemoteAddr() net.Addr               { return nil }
func (m *mockConn) SetDeadline(t time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }

func TestHandleConnection_ValidPacket(t *testing.T) {
	// Create a simulated session
	s := models.NewSession()
	ch := make(chan string, 1)
	s.PacketPreviews = list.New([]list.Item{}, list.NewDefaultDelegate(), 80, 24)
	s.Filter = models.NewFilter(s, 80, 24)

	// Fake packet that matches parsePacket() format
	fakePacket := "TCP;1234;10.0.0.1;10.0.0.2;80;1712587200.123456;64\n"
	conn := &mockConn{Reader: bytes.NewReader([]byte(fakePacket))}

	go handleConnection(conn, s, ch)

	time.Sleep(100 * time.Millisecond) // wait for goroutine

	if len(s.GetPackets()) == 0 {
		t.Fatal("expected session to receive one parsed packet")
	}

	p := s.GetPackets()[0]
	if p.SourceIP() != "10.0.0.1" || p.DestinationIP() != "10.0.0.2" {
		t.Errorf("unexpected IPs: %s -> %s", p.SourceIP(), p.DestinationIP())
	}
	if p.SourcePort() != "1234" || p.DestinationPort() != "80" {
		t.Errorf("unexpected ports: %s -> %s", p.SourcePort(), p.DestinationPort())
	}
	if p.Protocol() != "TCP" {
		t.Errorf("unexpected protocol: %s", p.Protocol())
	}
	if p.TTLValue() != 64 {
		t.Errorf("unexpected TTL: %d", p.TTLValue())
	}
}
