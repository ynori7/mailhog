package data

import (
	"testing"
)

func TestExtractBoundary(t *testing.T) {
	contents := []struct {
		content string
		expect  string
	}{
		{
			`multipart/alternative; boundary="_----------=_MCPart_498914860"`,
			`_----------=_MCPart_498914860`,
		},
		{
			`multipart/alternative; boundary=047d7bd74a2049b624050d805118`,
			`047d7bd74a2049b624050d805118`,
		},
	}

	for _, c := range contents {
		if b := extractBoundary(c.content); b != c.expect {
			t.Fatal("extractBoundary expect", c.expect, "but get", b)
		}
	}
}
