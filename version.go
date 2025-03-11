package wazuhapi

import (
	"encoding/json"
)

func (w *WazuhAPI) GetApiVersion() (version string, err error) {
	resp, err := w.Call("/", "GET", "")
	if err != nil {
		return "", err
	}

	type Response struct {
		Data struct {
			ApiVersion string `json:"api_version"`
		}
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return "", err
	}

	return response.Data.ApiVersion, nil
}
