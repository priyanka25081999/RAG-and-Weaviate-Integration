package document

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadToSharePoint(client *http.Client, filepath, sharepointURL, metadata string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	_, err = io.Copy(body, file)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", sharepointURL, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/pdf")
	req.Header.Set("Metadata", metadata)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	return nil
}
