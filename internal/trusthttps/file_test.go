package trusthttps

import "testing"

// TestFilesWatchStartStop tests Start() and Stop() of FilesWatch
func TestFilesWatchStartStop(t *testing.T) {
	probes := make(chan struct{})
	fw := NewFilesWatch(probes)
	fw.Start()
	fw.Stop()
}

// TestNewFilesWatch tests NewFilesWatch
func TestNewFilesWatch(t *testing.T) {
	probes := make(chan struct{})
	fw := NewFilesWatch(probes)
	if fw.probes != probes {
		t.Errorf("got %p, want %p", fw.probes, probes)
	}
	if fw.done == nil {
		t.Errorf("got nil, want != nil")
	}
	if fw.closed == nil {
		t.Errorf("got nil, want != nil")
	}
}
