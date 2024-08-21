package main

import (
	"fmt"

	"github.com/personal/RAG-and-Weaviate-Integration/internal/document"
	indexing "github.com/personal/RAG-and-Weaviate-Integration/internal/index"
	"github.com/personal/RAG-and-Weaviate-Integration/internal/query"
	"github.com/personal/RAG-and-Weaviate-Integration/internal/rag"
)

func main() {
	// Sample main logic integrating different components
	filepath := "financial-document.pdf"
	metadata := map[string]string{"Department": "Finance"}

	// Download the PDF document
	if err := document.DownloadFile("https://example.com/document.pdf", filepath); err != nil {
		fmt.Println("Failed to download file:", err)
		return
	}

	// Convert and Index the document
	text, err := document.ConvertPDFToText(filepath)
	if err != nil {
		fmt.Println("Failed to convert PDF:", err)
		return
	}
	if err := indexing.IndexDocumentInWeaviate(text, metadata); err != nil {
		fmt.Println("Failed to index document:", err)
		return
	}

	// Query and Generate Answer
	queryText := "show me documents related to finance"
	results, err := query.QueryWeaviate(queryText, "Finance")
	if err != nil {
		fmt.Println("Failed to query Weaviate:", err)
		return
	}
	answer := rag.GenerateAnswer(queryText, results)
	fmt.Println(answer)
}
