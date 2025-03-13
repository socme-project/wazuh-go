package main

import (
	"fmt"

	wazuhapi "github.com/socme-project/wazuh-go"
)

func main() {
	wazuh := wazuhapi.WazuhAPI{
		Host:     "10.8.178.20",
		Port:     "55000",
		Username: "admin",
		Password: "HMthisismys3cr3tP5ssword34a;",
		Indexer: wazuhapi.Indexer{
			Username: "admin",
			Password: "HMthisismys3cr3tP5ssword34a;",
			Host:     "10.8.178.20",
			Port:     "9200",
		},
		Insecure: true,
	}

	err := wazuh.RefreshToken()
	if err != nil {
		panic(err)
	}

	lastAlertId := 0
	alerts, lastAlertId, err := wazuh.GetAlerts(lastAlertId)
	if err != nil {
		panic(err)
	}
	// send alerts to the database
	fmt.Println(len(alerts))

	// every 5 mins getAlerts lastAlertId
	alerts, lastAlertId, err = wazuh.GetAlerts(lastAlertId)
	// debug
	fmt.Println(len(alerts))
}
