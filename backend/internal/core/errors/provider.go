package cErrors

var (
	ErrUnableToGetProvider = &APIError{
		Code:    "UNABLE_TO_GET_PROVIDER",
		Message: "We are unable to get provider details.",
	}

	ErrUnableToGetProviderCreds = &APIError{
		Code:    "UNABLE_TO_GET_PROVIDER_CREDS",
		Message: "We are unable to get provider credentials.",
	}
)
