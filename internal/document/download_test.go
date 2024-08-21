package document

import (
	"testing"
)

func TestDownloadFile(t *testing.T) {
	url := "https://example.com/document.pdf"
	filepath := "test-document.pdf"

	err := DownloadFile(url, filepath)
	if err != nil {
		t.Errorf("Failed to download file: %v", err)
	}
}
