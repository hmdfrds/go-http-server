package utils_test

import (
	"go-http-server/utils"
	"testing"
)

func TestUnescapeString(t *testing.T) {
	escaped := "Hello%20World%21"
	unescaped := utils.UnescapeString(escaped)
	expected := "Hello World!"

	if unescaped != expected {
		t.Errorf("Expected '%s', got '%s'", expected, unescaped)
	}
}
