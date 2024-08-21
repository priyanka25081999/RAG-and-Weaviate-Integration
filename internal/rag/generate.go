package rag

import (
	"fmt"
	"strings"

	"github.com/personal/RAG-and-Weaviate-Integration/internal/query"
)

func GenerateAnswer(queryText string, results []query.QueryResult) string {
	for _, result := range results {
		if strings.Contains(strings.ToLower(result.Document), strings.ToLower(queryText)) {
			return fmt.Sprintf("Answer: %s", result.Document)
		}
	}
	return "Sorry, I couldn't find a relevant document."
}
