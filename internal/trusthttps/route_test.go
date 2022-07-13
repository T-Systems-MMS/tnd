package trusthttps

import "testing"

// TestRoutesWatchStartStop tests Start and Stop of RoutesWatch
func TestRoutesWatchStartStop(t *testing.T) {
	probes := make(chan struct{})
	rw := NewRoutesWatch(probes)
	rw.Start()
	rw.Stop()
}

// TestRoutesWatchProbes tests Probes of RoutesWatch
func TestRoutesWatchProbes(t *testing.T) {
	probes := make(chan struct{})
	rw := NewRoutesWatch(probes)
	if rw.Probes() != probes {
		t.Errorf("got %p, want %p", rw.Probes(), probes)
	}
}

// TestNewRoutesWatch tests NewRoutesWatch
func TestNewRoutesWatch(t *testing.T) {
	probes := make(chan struct{})
	rw := NewRoutesWatch(probes)
	if rw.probes != probes {
		t.Errorf("got %p, want %p", rw.probes, probes)
	}
	if rw.done == nil {
		t.Errorf("got nil, want != nil")
	}
}
