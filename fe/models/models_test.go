package models

import (
	"testing"
	"time"
)

// --- Validation Tests ---

func TestValidateIP(t *testing.T) {
	valid := []string{"192.168.1.1", "0.0.0.0", "255.255.255.255"}
	invalid := []string{"256.256.256.256", "192.168.1", "abc.def.ghi.jkl", "1234"}

	for _, ip := range valid {
		if err := ValidateIP(ip); err != nil {
			t.Errorf("expected valid IP '%s', got error: %v", ip, err)
		}
	}

	for _, ip := range invalid {
		if err := ValidateIP(ip); err == nil {
			t.Errorf("expected invalid IP '%s', got no error", ip)
		}
	}
}

func TestValidatePort(t *testing.T) {
	valid := []string{"80", "443", "0", "9999"}
	invalid := []string{"-1", "65536", "abcd", "12345"}

	for _, port := range valid {
		if err := ValidatePort(port); err != nil {
			t.Errorf("expected valid port '%s', got error: %v", port, err)
		}
	}

	for _, port := range invalid {
		if err := ValidatePort(port); err == nil {
			t.Errorf("expected invalid port '%s', got no error", port)
		}
	}
}

func TestValidateTime(t *testing.T) {
	valid := "2023-04-08 12:00:00 "
	invalid := "2023-04-08T12:00:00"

	if err := ValidateTime(valid); err != nil {
		t.Errorf("expected valid time, got error: %v", err)
	}
	if err := ValidateTime(invalid); err == nil {
		t.Error("expected invalid time format, got no error")
	}
}

// --- Filter Tests ---

func makePacket(ip, srcPort, dstPort, proto string, timestamp time.Time) Packet {
	network := Network{
		SourceIP: ip,
	}
	transport := Transport{
		SourcePort:      srcPort,
		DestinationPort: dstPort,
		Protocol:        proto,
	}
	return Packet{
		Network:   network,
		Transport: transport,
		datetime:  timestamp,
	}
}

func TestFilter_Matches_Fields(t *testing.T) {
	now := time.Now()

	packet := makePacket("192.168.0.1", "1234", "80", "TCP", now)

	filter := &Filter{
		IP:                "192.168.0.1",
		senderPort:        "1234",
		receiverPort:      "80",
		transportProtocol: "TCP",
		startTime:         now.Add(-time.Minute),
		endTime:           now.Add(time.Minute),
		active:            true,
	}

	if !filter.Matches(packet) {
		t.Error("expected full match on all fields")
	}

	// Change IP
	filter.IP = "10.0.0.1"
	if filter.Matches(packet) {
		t.Error("expected IP mismatch to fail")
	}
	filter.IP = "192.168.0.1"

	// Change sender port
	filter.senderPort = "9999"
	if filter.Matches(packet) {
		t.Error("expected sender port mismatch to fail")
	}
	filter.senderPort = "1234"

	// Change receiver port
	filter.receiverPort = "9999"
	if filter.Matches(packet) {
		t.Error("expected receiver port mismatch to fail")
	}
	filter.receiverPort = "80"

	// Change protocol
	filter.transportProtocol = "UDP"
	if filter.Matches(packet) {
		t.Error("expected protocol mismatch to fail")
	}
	filter.transportProtocol = "TCP"

	// Out of time range
	filter.startTime = now.Add(time.Second)
	if filter.Matches(packet) {
		t.Error("expected start time cutoff to fail")
	}
	filter.startTime = now.Add(-time.Minute)
	filter.endTime = now.Add(-time.Second)
	if filter.Matches(packet) {
		t.Error("expected end time cutoff to fail")
	}
}

func TestFilter_Reset(t *testing.T) {
	filter := &Filter{
		IP:                "1.2.3.4",
		senderPort:        "1111",
		receiverPort:      "2222",
		transportProtocol: "TCP",
		startTime:         time.Now(),
		endTime:           time.Now().Add(time.Minute),
		active:            true,
		width:             80,
		height:            24,
	}

	filter.Reset()

	if filter.IP != "" || filter.senderPort != "" || filter.receiverPort != "" || filter.transportProtocol != "" {
		t.Error("expected all fields to be reset")
	}

	if !filter.startTime.IsZero() || !filter.endTime.IsZero() || filter.active {
		t.Error("expected time range and active to be reset")
	}
}

func TestPacket_Getters(t *testing.T) {
	now := time.Now()
	network := Network{
		SourceIP: "10.0.0.1",
	}
	transport := Transport{
		SourcePort:      "1234",
		DestinationPort: "80",
		Protocol:        "TCP",
	}
	p := Packet{
		Network:   network,
		Transport: transport,
		datetime:  now,
	}

	if p.SourceIP() != "10.0.0.1" {
		t.Errorf("unexpected SourceIP: %s", p.SourceIP())
	}
	if p.SourcePort() != "1234" {
		t.Errorf("unexpected SenderPort: %s", p.SourcePort())
	}
	if p.DestinationPort() != "80" {
		t.Errorf("unexpected ReceiverPort: %s", p.DestinationPort())
	}
	if p.Protocol() != "TCP" {
		t.Errorf("unexpected Protocol: %s", p.Protocol())
	}
	if !p.Datetime().Equal(now) {
		t.Errorf("unexpected Timestamp")
	}
}

// --- PacketPreview Tests ---

func TestNewPacketPreview(t *testing.T) {
	network := Network{
		SourceIP: "192.168.1.1",
	}
	transport := Transport{
		Protocol: "TCP",
	}
	p := Packet{Transport: transport, Network: network}
	pp := NewPacketPreview(p)

	if pp.Title() == "" {
		t.Error("Title should not be empty")
	}
	if pp.Description() == "" {
		t.Error("PreviewContent should not be empty")
	}
}
