package trusthttps

import (
	"crypto/sha256"
	"encoding/hex"
	"net"
	"testing"
)

// TestTrustedHTTPSAddServer tests AddServer of TrustedHTTPS
func TestTrustedHTTPSAddServer(t *testing.T) {
	th := NewTrustedHTTPS()
	url := "http://test.example.com:442"
	cert := []byte("raw test certificate")
	hash := sha256.Sum256(cert)
	hs := hex.EncodeToString(hash[:])
	th.AddServer(url, hs)

	want := url
	got := th.servers[0].url
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	want = hs
	got = th.servers[0].hash
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestTrustedHTTPSSetDialer tests SetDialer of TrustedHTTPS
func TestTrustedHTTPSSetDialer(t *testing.T) {
	th := NewTrustedHTTPS()
	dialer := &net.Dialer{}
	th.SetDialer(dialer)

	want := dialer
	got := th.dialer
	if got != want {
		t.Errorf("got %p, want %p", got, want)
	}
}

// TestTrustedHTTPSProbe tests Probe of TrustedHTTPS
func TestTrustedHTTPSProbe(t *testing.T) {
	th := NewTrustedHTTPS()
	th.Start()
	th.Probe()
	th.Stop()
}

// TestTrustedHTTPSStartStop tests Start and Stop of TrustedHTTPS
func TestTrustedHTTPSStartStop(t *testing.T) {
	th := NewTrustedHTTPS()
	th.Start()
	th.Stop()
}

// TestTrustedHTTPSProbes tests Probes of TrustedHTTPS
func TestTrustedHTTPSProbes(t *testing.T) {
	th := NewTrustedHTTPS()
	want := th.probes
	got := th.Probes()
	if want != got {
		t.Errorf("got %p, want %p", got, want)
	}
}

// TestTrustedHTTPSResults tests Results of TrustedHTTPS
func TestTrustedHTTPSResults(t *testing.T) {
	th := NewTrustedHTTPS()
	want := th.results
	got := th.Results()
	if want != got {
		t.Errorf("got %p, want %p", got, want)
	}
}

// TestNewTrustedHTTPS tests NewTrustedHTTPS
func TestNewTrustedHTTPS(t *testing.T) {
	th := NewTrustedHTTPS()
	if th.probes == nil {
		t.Errorf("got nil, want != nil")
	}
	if th.results == nil {
		t.Errorf("got nil, want != nil")
	}
	if th.done == nil {
		t.Errorf("got nil, want != nil")
	}
	if th.dialer == nil {
		t.Errorf("got nil, want != nil")
	}
}
