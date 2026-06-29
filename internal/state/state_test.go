package state

import (
	"os"
	"testing"
	"time"
)

func TestMetadataRoundTrip(t *testing.T) {
	t.Setenv("HIVE_STATE_DIR", t.TempDir())

	meta := Metadata{
		Name:      "work",
		Tool:      "nvim",
		Path:      "/tmp/project",
		CreatedAt: time.Now().UTC().Truncate(time.Second),
	}

	if err := WriteMetadata(meta); err != nil {
		t.Fatal(err)
	}

	got, err := ReadMetadata("work")
	if err != nil {
		t.Fatal(err)
	}

	if got != meta {
		t.Fatalf("ReadMetadata() = %#v, want %#v", got, meta)
	}
}

func TestRuntimeRoundTrip(t *testing.T) {
	t.Setenv("HIVE_STATE_DIR", t.TempDir())

	endedAt := time.Now().UTC().Truncate(time.Second)
	exitCode := 0
	runtime := Runtime{
		State:     StateCompleted,
		StartedAt: endedAt.Add(-time.Minute),
		EndedAt:   &endedAt,
		ExitCode:  &exitCode,
	}

	if err := WriteRuntime("work", runtime); err != nil {
		t.Fatal(err)
	}

	got, err := ReadRuntime("work")
	if err != nil {
		t.Fatal(err)
	}

	if got.State != runtime.State || !got.StartedAt.Equal(runtime.StartedAt) {
		t.Fatalf("ReadRuntime() = %#v, want %#v", got, runtime)
	}

	if got.EndedAt == nil || !got.EndedAt.Equal(*runtime.EndedAt) {
		t.Fatalf("EndedAt = %#v, want %#v", got.EndedAt, runtime.EndedAt)
	}

	if got.ExitCode == nil || *got.ExitCode != *runtime.ExitCode {
		t.Fatalf("ExitCode = %#v, want %#v", got.ExitCode, runtime.ExitCode)
	}
}

func TestReadMissingRuntimeReturnsUnknown(t *testing.T) {
	t.Setenv("HIVE_STATE_DIR", t.TempDir())

	got, err := ReadRuntime("missing")
	if err != nil {
		t.Fatal(err)
	}

	if got.State != StateUnknown {
		t.Fatalf("State = %q, want %q", got.State, StateUnknown)
	}
}

func TestDeleteSessionState(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("HIVE_STATE_DIR", dir)

	if err := WriteMetadata(Metadata{Name: "work"}); err != nil {
		t.Fatal(err)
	}

	if err := DeleteSessionState("work"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(sessionDir("work")); !os.IsNotExist(err) {
		t.Fatalf("session dir still exists or unexpected error: %v", err)
	}
}
