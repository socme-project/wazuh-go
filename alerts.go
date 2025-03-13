package wazuhapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Alert struct {
	WazuhAlertID    string `json:"wazuh_alert_id"`
	RuleID          string `json:"rule_id"`
	RuleLevel       uint   `json:"rule_level"`
	RuleDescription string `json:"rule_description"`
	Timestamp       string `json:"timestamp"`
	RawJSON         string `json:"raw_json"`
	Sort            int    `json:"sort"`
}

func (w *WazuhAPI) GetAlerts(lastAlertId int) (alerts []Alert, lastId int, err error) {
	type Response struct {
		Hits struct {
			Hits []struct {
				ID     string `json:"_id"`
				Source struct {
					Rule struct {
						ID          string
						Level       uint
						Description string
					}
					Data struct {
						Description string
					}
					Timestamp string
				} `json:"_source"`
				Sort []int
			}
		}
	}

	alerts = []Alert{}

	for {

		// fmt.Println("Starting")
		var query string
		if lastAlertId != 0 {
			query = `{ "size": 500, "sort": [ { "timestamp": { "order": "asc" } } ], "search_after": [` + strconv.Itoa(
				lastAlertId,
			) + `]}`
		} else {
			query = `{ "size": 500, "sort": [ { "timestamp": { "order": "asc" } } ] }`
		}
		req, err := http.NewRequest(
			"GET",
			"https://"+w.Indexer.Host+":"+w.Indexer.Port+"/wazuh-alerts-*/_search/",
			bytes.NewBuffer([]byte(query)),
		)
		if err != nil {
			return nil, lastAlertId, err
		}

		req.SetBasicAuth(w.Indexer.Username, w.Indexer.Password)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, lastAlertId, err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, lastAlertId, err
		}

		var response Response

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, lastAlertId, err
		}
		// fmt.Println("Response: ", response)

		for _, item := range response.Hits.Hits {
			raw, err := json.Marshal(item)
			if err != nil {
				return nil, lastAlertId, err
			}

			if item.Source.Rule.Description == "" {
				if item.Source.Data.Description == "" {
					item.Source.Rule.Description = "No title and description available"
				}
				item.Source.Rule.Description = item.Source.Data.Description
			}
			alerts = append(alerts, Alert{
				WazuhAlertID:    item.ID,
				RuleID:          item.Source.Rule.ID,
				RuleLevel:       item.Source.Rule.Level,
				RuleDescription: item.Source.Rule.Description,
				Timestamp:       item.Source.Timestamp,
				// TODO: FIX RAW
				RawJSON:         string(raw),
				Sort:            item.Sort[0],
			})

		}
		if len(response.Hits.Hits) == 0 {
			return alerts, lastAlertId, nil
		}
		lastAlertId = alerts[len(alerts)-1].Sort

		if len(response.Hits.Hits) < 500 {
			return alerts, lastAlertId, nil
		}
	}
}
