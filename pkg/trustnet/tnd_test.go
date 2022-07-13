package trustnet

import (
	"crypto/sha256"
	"encoding/hex"
	"net"
	"testing"
)

// TestTNDAddServer tests AddServer of TND
func TestTNDAddServer(t *testing.T) {
	tnd := NewTND()
	url := "http://test.example.com:442"
	cert := []byte("raw test certificate")
	hash := sha256.Sum256(cert)
	hs := hex.EncodeToString(hash[:])
	tnd.AddServer(url, hs)
}

// TestTNDSetDialer tests SetDialer of TND
func TestTNDSetDialer(t *testing.T) {
	tnd := NewTND()
	dialer := &net.Dialer{}
	tnd.SetDialer(dialer)
}

// TestTNDStartStop tests Start and Stop of TND
func TestTNDStartStop(t *testing.T) {
	tnd := NewTND()
	tnd.Start()
	tnd.Stop()
}

// TestTNDResults tests Results of TND
func TestTNDResults(t *testing.T) {
	tnd := NewTND()
	results := tnd.Results()
	if results == nil {
		t.Errorf("got nil, want != nil")
	}
}

// TestNewTND tests NewTND
func TestNewTND(t *testing.T) {
	tnd := NewTND()
	if tnd.https == nil {
		t.Errorf("got nil, want != nil")
	}
	if tnd.results == nil {
		t.Errorf("got nil, want != nil")
	}
	if tnd.done == nil {
		t.Errorf("got nil, want != nil")
	}
}
