package cErrors

var (
	ErrUserAlreadyExists = &APIError{
		Code:    "USER_EXISTS",
		Message: "User with this email already exists.",
	}

	ErrUserDoesNotExists = &APIError{
		Code:    "USER_NOT_EXISTS",
		Message: "User with this email does not exists.",
	}

	ErrWrongPassword = &APIError{
		Code:    "WRONG_PASSWORD",
		Message: "Wrong Password. Try Again.",
	}

	ErrUnableToCreateAccount = &APIError{
		Code:    "UNABLE_TO_CREATE_ACCOUNT",
		Message: "Unable to create account. Please try again later.",
	}

	ErrUnableToLogin = &APIError{
		Code:    "UNABLE_TO_LOGIN",
		Message: "Unable to login. Please try again later.",
	}

	ErrUnableToCompleteOnboarding = &APIError{
		Code:    "UNABLE_TO_COMPLETE_ONBOARDING",
		Message: "Unable to complete onboarding. Please try again later.",
	}
)
