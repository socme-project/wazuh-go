package wazuhapi

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
)

// Call is a method of WazuhAPI that sends a request to the Wazuh API
// and returns the response or an error.
// Additionaly, it can handle basic auth by adding "+BA" to the method.
func (w WazuhAPI) Call(path, method, json string) (response []byte, err error) {
	if w.Insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// Check if the method needs basic auth
	basicAuth := false
	if strings.HasSuffix(method, "+BA") {
		method = strings.TrimSuffix(method, "+BA")
		basicAuth = true
	}

	req, err := http.NewRequest(
		method,
		"https://"+w.Host+":"+w.Port+path,
		bytes.NewBuffer([]byte(json)),
	)
	if err != nil {
		return nil, err
	}
	if basicAuth {
		req.SetBasicAuth(w.Username, w.Password)
	} else {
		req.Header.Add("Authorization", "Bearer "+w.Token)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
