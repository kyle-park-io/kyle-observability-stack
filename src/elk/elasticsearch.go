package elk

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

func NewClient() (*elasticsearch.Client, error) {

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
		return nil, err
	}

	// Check Elasticsearch status
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	fmt.Println("Elasticsearch Info:")
	fmt.Println(res)

	return es, nil
}
