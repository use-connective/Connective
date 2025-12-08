package connectors

import (
	"strings"
)

func maskCode(code string) string {
	if code == "" || len(code) <= 8 {
		return "***"
	}
	return code[:4] + "***" + code[len(code)-4:]
}

func maskState(state string) string {
	if state == "" || len(state) <= 10 {
		return "***"
	}
	return state[:5] + "***" + state[len(state)-5:]
}

func maskClientID(clientID string) string {
	if clientID == "" || len(clientID) <= 8 {
		return "***"
	}
	return clientID[:4] + "***" + clientID[len(clientID)-4:]
}

func checkForGoogleAuth(providerName string, redirectURL string) string {
	if strings.HasPrefix(providerName, "google") {
		redirectURL = redirectURL + "&access_type=offline&prompt=consent&response_type=code"
	}

	return redirectURL
}
