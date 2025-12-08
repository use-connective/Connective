package iSlack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/x-sushant-x/connective/internal/connectors/common"
)

func (s *SlackConnector) SendMessage(ctx context.Context, creds *common.UserCredentials, payload map[string]any) (any, error) {

	bodyMap := map[string]any{
		"channel": payload["channel"],
		"text":    payload["text"],
	}

	body, _ := json.Marshal(bodyMap)

	req, _ := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+creds.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out map[string]any
	json.NewDecoder(resp.Body).Decode(&out)

	return out, nil
}

func (s *SlackConnector) ListChannels(ctx context.Context, creds *common.UserCredentials, payload map[string]any) (any, error) {

	req, _ := http.NewRequest("GET", "https://slack.com/api/conversations.list", nil)
	req.Header.Set("Authorization", "Bearer "+creds.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out map[string]any
	json.NewDecoder(resp.Body).Decode(&out)

	return out, nil
}
