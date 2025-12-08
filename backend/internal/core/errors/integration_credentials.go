package cErrors

var (
	ErrUnableToSaveCredentials = &APIError{
		Code:    "UNABLE_TO_SAVE_INTEGRATION_CREDS",
		Message: "Unable to save provider credentials. Please try again later.",
	}

	ErrUnableToGetProvidersList = &APIError{
		Code:    "UNABLE_TO_GET_PROVIDERS_LIST",
		Message: "Unable to get providers list. Please try again later.",
	}

	ErrUnableToGetConnectedAccount = &APIError{
		Code:    "UNABLE_TO_GET_CONNECTED_ACCOUNT",
		Message: "Unable to get connected accounts. Please try again later.",
	}
)
