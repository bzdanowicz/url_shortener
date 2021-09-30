package utils

import (
	"testing"
)

func TestEncodeUrl(t *testing.T) {
	const exampleUrl = "http://google.com"

	encodedUrl_1, _ := EncodeUrl(exampleUrl)
	encodedUrl_2, _ := EncodeUrl(exampleUrl)

	if encodedUrl_1 != encodedUrl_2 {
		t.Errorf("Expected encoded urls to be the same")
	}
}
