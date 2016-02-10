package libsass

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	var (
		expect  = "3.2.5"
		version = Version()
	)

	if !strings.HasPrefix(version, expect) {
		t.Fatalf("Expected Version: %s, Got: %s", expect, version)
	}

	t.Logf("Version %s", version)

}
