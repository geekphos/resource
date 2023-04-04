package errno

var (
	// ErrResourceAlreadyExists resource already exists.
	ErrResourceAlreadyExists = &Errno{HTTP: 400, Code: "FailedOperation.ErrResourceAlreadyExist", Message: "Menu already exist."}

	// ErrResourceNotFound resource not found.
	ErrResourceNotFound = &Errno{HTTP: 404, Code: "FailedOperation.ErrResourceNotFound", Message: "Menu not found."}
)
