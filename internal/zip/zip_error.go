package zip

import "fmt"

const defaultMessage = "Error creating zip file"

type ZipError struct {
	Err     error
	Context string
}

func (ze *ZipError) Unwrap() error { return ze.Err }

func (ze *ZipError) ErrorMessage() string {
	if len(ze.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, ze.Context, ze.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, ze.Err)
	}
}
