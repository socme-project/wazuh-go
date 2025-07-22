# Wazuh Go Client

This project is a Go client library designed to interact with the Wazuh API and
retrieve alerts directly from the Wazuh Indexer (Elasticsearch/OpenSearch).

## Features

- **Authentication:** Handles authentication and refreshing of Wazuh API tokens.
- **API Version:** Allows retrieving the Wazuh API version.
- **Agent Status:** Provides information on the status of agents (active,
  disconnected, total, synced, not synced).
- **Alert Retrieval:** Fetches alerts from the Wazuh Indexer with pagination
  support.

---

## Installation

To use this library in your Go project, run the following command:

```bash
go get github.com/socme-project/wazuh-go
```

---

## Usage

Here's an example of how to use the library to refresh a token and retrieve
alerts:

```go
package main

import (
	"fmt"
	"os"

	wazuhapi "github.com/socme-project/wazuh-go"
)

func main() {
	// Initialize the WazuhAPI struct with your connection details.
	// It is highly recommended to use environment variables
	// or a secret management system for credentials in production.
	wazuh := wazuhapi.WazuhAPI{
		Host:     "10.8.178.20", // Replace with your Wazuh API IP/hostname
		Port:     "55000",       // Wazuh API port
		Username: "admin",
		Password: "HMthisismys3cr3tP5ssword34a;", // Wazuh API user password
		Indexer: wazuhapi.Indexer{
			Username: "admin",
			Password: "HMthisismys3cr3tP5ssword34a;", // Indexer user password
			Host:     "10.8.178.20",                   // Replace with your Indexer IP/hostname
			Port:     "9200",                          // Indexer port
		},
		Insecure: true, // true to skip TLS certificate verification (not recommended in production)
	}

	// Refresh the authentication token
	err := wazuh.RefreshToken()
	if err != nil {
		fmt.Printf("Error refreshing token: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Wazuh API token refreshed successfully.")

	// Retrieve alerts
	lastAlertId := 0 // Use 0 for the first request, then the last alert ID for pagination.
	alerts, newLastAlertId, err := wazuh.GetAlerts(lastAlertId)
	if err != nil {
		fmt.Printf("Error retrieving alerts: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of alerts retrieved: %d\n", len(alerts))
	fmt.Printf("Last alert ID for the next request: %d\n", newLastAlertId)

	// Example of retrieving agent status
	agentsStatus, err := wazuh.GetAgents()
	if err != nil {
		fmt.Printf("Error retrieving agent status: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Agent Status: Active=%d, Disconnected=%d, Total=%d\n",
		agentsStatus.Active, agentsStatus.Disconnected, agentsStatus.Total)

	// Example of retrieving API version
	apiVersion, err := wazuh.GetApiVersion()
	if err != nil {
		fmt.Printf("Error retrieving API version: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wazuh API Version: %s\n", apiVersion)
}
```

---

## Contributing

Contributions are welcome ! Feel free to open an issue or submit a Pull Request.
