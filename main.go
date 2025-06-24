package wazuhapi

import (
	"encoding/json"
)

func (w *WazuhAPI) RefreshToken() (err error) {
	resp, err := w.Call("/security/user/authenticate", "POST+BA", "")
	if err != nil {
		return err
	}

	type Response struct {
		Data struct {
			Token string `json:"token"`
		}
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	if err != nil || response.Data.Token == "" {
		return err
	}
	w.Token = response.Data.Token

	return nil
}
