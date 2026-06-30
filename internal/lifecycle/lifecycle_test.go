package lifecycle

import "testing"

func TestIsBuiltinToolRecognizesSupportedLaunchProfiles(t *testing.T) {
	for _, tool := range []string{"pi", "claude", "nvim", "bash"} {
		if !IsBuiltinTool(tool) {
			t.Fatalf("expected %q to be builtin", tool)
		}
	}

	if IsBuiltinTool("cursor") {
		t.Fatal("expected unknown tool to be rejected")
	}
}

func TestCreateRejectsUnknownToolBeforeTouchingTmux(t *testing.T) {
	svc := Service{
		generateID: func(int) string { return "fixed" },
		workingDir: func() (string, error) { return "/tmp", nil },
	}

	_, err := svc.Create(CreateRequest{Tool: "cursor"})
	if err == nil {
		t.Fatal("expected unknown tool error")
	}
}

func TestResolvePathUsesWorkingDirForEmptyOrDot(t *testing.T) {
	svc := Service{workingDir: func() (string, error) { return "/work/project", nil }}

	for _, value := range []string{"", "."} {
		got, err := svc.resolvePath(value)
		if err != nil {
			t.Fatal(err)
		}
		if got != "/work/project" {
			t.Fatalf("resolvePath(%q) = %q", value, got)
		}
	}
}
