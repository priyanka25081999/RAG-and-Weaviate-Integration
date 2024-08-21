package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type QueryResult struct {
	Document string `json:"document"`
}

func QueryWeaviate(query, department string) ([]QueryResult, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/v1/query?query=%s", query), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("WEAVIATE_API_KEY"))
	req.Header.Set("Department", department)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []QueryResult
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
