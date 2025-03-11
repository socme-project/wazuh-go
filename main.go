package wazuhapi

import (
	"encoding/json"
)

// curl -u <USER>:<PASSWORD> -k -X POST "https://<HOST_IP>:55000/security/user/authenticate"
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
