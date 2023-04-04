package errno

var (
	// ErrMenuAlreadyExist is used when the menu already exists.
	ErrMenuAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.ErrMenuAlreadyExist", Message: "Menu already exist."}

	// ErrMenuNotFound is used when the menu is not found.
	ErrMenuNotFound = &Errno{HTTP: 404, Code: "FailedOperation.ErrMenuNotFound", Message: "Menu not found."}
)
