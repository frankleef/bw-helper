package vault

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/frankleef/bw-login/internal/config"
)

func Unlock() (*string, error) {
	postBody, _ := json.Marshal(map[string]string{
		"password": config.Configuration.Password,
	})

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(fmt.Sprintf("%s://%s:%d/unlock", config.Configuration.Scheme, config.Configuration.Host, config.Configuration.Port), "application/json", responseBody)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	json.Unmarshal(body, &result)
	if !result.Success {
		return nil, errors.New("unexpected response from Bitwarden, success is false")
	}

	return &result.Data.Raw, nil
}
