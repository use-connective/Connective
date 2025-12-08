package cErrors

var (
	ErrProjectExists = &APIError{
		Code:    "PROJECT_EXISTS",
		Message: "Project already exists",
	}

	ErrUnableToCreateProject = &APIError{
		Code:    "UNABLE_TO_CREATE_PROJECT",
		Message: "Unable to create project",
	}
)
