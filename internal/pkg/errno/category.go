package errno

var (
	// ErrCategoryAlreadyExist is returned when the category already exists.
	ErrCategoryAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.CategoryAlreadyExist", Message: "Category already exist."}

	// ErrCategoryNotFound is returned when the category is not found.
	ErrCategoryNotFound = &Errno{HTTP: 404, Code: "FailedOperation.CategoryNotFound", Message: "Category not found."}
)
