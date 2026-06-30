package tui

import "testing"

func TestAppStatePreventsStackedOverlays(t *testing.T) {
	state := NewAppState()

	state.ShowHelp()
	state.Search()

	if !state.Searching() {
		t.Fatal("expected search overlay to replace help overlay")
	}
	if state.ShowingHelp() {
		t.Fatal("expected help overlay to be closed")
	}
}

func TestAppStateReturningToSessionListClosesOverlay(t *testing.T) {
	state := NewAppState()
	state.StartNewSession()
	state.PickPath()

	state.ReturnToSessionList()

	if state.Screen != ScreenList {
		t.Fatalf("expected list screen, got %q", state.Screen)
	}
	if state.Overlay != OverlayNone {
		t.Fatalf("expected no overlay, got %q", state.Overlay)
	}
}
