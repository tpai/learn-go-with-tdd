package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "wazup://whats.up" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://wtfismyip.com",
		"wazup://whats.up",
	}

	expected := map[string]bool{
		"https://google.com":    true,
		"https://wtfismyip.com": true,
		"wazup://whats.up":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected '%v' but got '%v'", expected, got)
	}
}
