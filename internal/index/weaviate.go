package indexing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeaviateDocument struct {
	Text     string            `json:"text"`
	Metadata map[string]string `json:"metadata"`
}

func IndexDocumentInWeaviate(text string, metadata map[string]string) error {
	document := WeaviateDocument{
		Text:     text,
		Metadata: metadata,
	}

	jsonData, err := json.Marshal(document)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/documents", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("WEAVIATE_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to index document: %s", resp.Status)
	}

	return nil
}
