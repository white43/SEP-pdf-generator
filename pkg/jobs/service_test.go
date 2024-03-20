package jobs

import (
	"strings"
	"testing"
)

func TestHTMLPayloadLength(t *testing.T) {
	service := NewDummyService()

	if service.IsHTMLPayloadLengthValid(strings.Repeat("a", 255*1024-1)) != true {
		t.Errorf("payloads less than 255 KB should pass")
	}

	if service.IsHTMLPayloadLengthValid(strings.Repeat("a", 255*1024)) != true {
		t.Errorf("payloads equal 255 KB should pass")
	}

	if service.IsHTMLPayloadLengthValid(strings.Repeat("a", 255*1024+1)) != false {
		t.Errorf("payloads more than 255 KB should not pass")
	}
}

func TestPrintableCharacters(t *testing.T) {
	service := NewDummyService()

	if service.DoesStringContainOnlyPrintableCharacters("aaa") != true {
		t.Errorf("payloads containing printable characters should pass")
	}

	if service.DoesStringContainOnlyPrintableCharacters("a\x00a") != false {
		t.Errorf("payloads containing non-printable characters should not pass")
	}
}
