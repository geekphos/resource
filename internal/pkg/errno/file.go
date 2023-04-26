package errno

var (
	// ErrFileNotFound
	ErrFileNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.FileNotFound", Message: "File not found."}
)
