package elk

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// go test -v -run TestGetJSON
func TestGetJSON(t *testing.T) {

	es, err := NewClient()
	if err != nil {
		t.Log(err)
	}

	// Index and Document ID to retrieve
	indexName := "my-index" // Specify the index name
	docID := "1"            // Specify the document ID

	// Create the GET request to retrieve the document
	req := esapi.GetRequest{
		Index:      indexName, // Index name
		DocumentID: docID,     // Document ID
	}

	// Execute the GET request
	res, err := req.Do(context.Background(), es)
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

// go test -v -run TestIndexJSON
func TestIndexJSON(t *testing.T) {

	es, err := NewClient()
	if err != nil {
		t.Log(err)
	}

	// Open the JSON file
	jsonFile, err := os.Open("json/data.json")
	if err != nil {
		log.Fatalf("Error opening JSON file: %s", err)
	}
	defer jsonFile.Close()

	// Read the contents of the JSON file using io.ReadAll
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	// Index the JSON data into Elasticsearch
	indexName := "my-index"
	docID := "1"
	req := esapi.IndexRequest{
		Index:      indexName,                  // Specify the index name
		DocumentID: docID,                      // Specify the document ID
		Body:       bytes.NewReader(byteValue), // Attach the JSON data
		Refresh:    "true",                     // Refresh immediately after indexing
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()

	// Print the indexing response
	fmt.Println("Index Response:")
	fmt.Println(res)
}

// go test -v -run TestNewClient
func TestNewClient(t *testing.T) {

	// Initialize the Elasticsearch client
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"}, // Access via HTTPS
		Username:  "elastic",                          // Username
		Password:  "xr2Y2xd4qVQX_rDdpp-n",             // Initial password (Input your password)
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Disable TLS certificate verification
			},
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Check Elasticsearch status
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Elasticsearch Info:")
	fmt.Println(res)
}
