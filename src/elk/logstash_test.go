package elk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// go test -v -run TestParseAndIndexLog
func TestParseAndIndexLog(t *testing.T) {

	es, err := NewClient()
	if err != nil {
		t.Log(err)
	}

	// Example log line
	logLine := `127.0.0.1 - - [17/Jan/2025:11:58:53 +0000] "GET /index.html HTTP/1.1" 200 1043`

	// Parse the log line
	parsedLog, err := parseLogLine(logLine)
	if err != nil {
		log.Fatalf("Error parsing log line: %s", err)
	}
	fmt.Println("parsedLog: ", parsedLog)

	// Convert to JSON
	jsonData, err := json.Marshal(parsedLog)
	if err != nil {
		log.Fatalf("Error converting to JSON: %s", err)
	}
	fmt.Println("bytes: ", jsonData)

	// Index the data in Elasticsearch
	indexName := "apache-logs"
	docID := "1" // Replace with a unique ID generator if needed
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: docID,
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

	fmt.Println("Indexed Document:", res.Status())

	// Create the GET request to retrieve the document
	req2 := esapi.GetRequest{
		Index:      indexName, // Index name
		DocumentID: docID,     // Document ID
	}

	// Execute the GET request
	res, err = req2.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error retrieving document: %s", err)
	}
	defer res.Body.Close()

	// Check for errors in the response
	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.Status())
	}

	// Print the retrieved document
	fmt.Println("Retrieved Document:")
	fmt.Println(res)
}

// go test -v -run TestParseLog
func TestParseLog(t *testing.T) {

	// Example log line
	logLine := `127.0.0.1 - - [17/Jan/2025:11:58:53 +0000] "GET /index.html HTTP/1.1" 200 1043`

	// Parse the log line
	parsedLog, err := parseLogLine(logLine)
	if err != nil {
		log.Fatalf("Error parsing log line: %s", err)
	}
	fmt.Println("parsedLog: ", parsedLog)

	// Convert to JSON
	jsonData, err := json.Marshal(parsedLog)
	if err != nil {
		log.Fatalf("Error converting to JSON: %s", err)
	}
	fmt.Println("bytes: ", jsonData)
}
