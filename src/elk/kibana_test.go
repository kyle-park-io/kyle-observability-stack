package elk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// go test -v -run TestIndexUserActivity
func TestIndexUserActivity(t *testing.T) {

	es, err := NewClient()
	if err != nil {
		t.Log(err)
	}

	// Example data to index
	data := []map[string]interface{}{
		{"timestamp": time.Now().Format(time.RFC3339), "user": "John Doe", "action": "login", "status": "success"},
		{"timestamp": time.Now().Format(time.RFC3339), "user": "Jane Smith", "action": "logout", "status": "success"},
		{"timestamp": time.Now().Format(time.RFC3339), "user": "Alice", "action": "login", "status": "failure"},
	}

	// Index each document
	indexName := "user-activity"
	for i, doc := range data {
		jsonData, err := json.Marshal(doc)
		if err != nil {
			log.Fatalf("Error marshalling document: %s", err)
		}

		req := esapi.IndexRequest{
			Index:      indexName,
			DocumentID: fmt.Sprintf("%d", i+1),
			Body:       bytes.NewReader(jsonData),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error indexing document: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Fatalf("Error response from Elasticsearch: %s", res.Status())
		}

		fmt.Printf("Indexed document %d: %s\n", i+1, res.Status())
	}

	fmt.Println("Data indexed successfully. You can now visualize it in Kibana.")
}
