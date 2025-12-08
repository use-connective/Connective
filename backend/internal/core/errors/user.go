package cErrors

var (
	ErrUnableToUpdateUser = &APIError{
		Code:    "UNABLE_TO_UPDATE_USER",
		Message: "Unable to update user. Please try again later.",
	}

	ErrUnableToDeleteUser = &APIError{
		Code:    "UNABLE_TO_DELETE_USER",
		Message: "Unable to delete user. Please try again later.",
	}

	ErrUnableToGetUser = &APIError{
		Code:    "UNABLE_TO_GET_USER",
		Message: "Unable to get user. Please try again later.",
	}
)
