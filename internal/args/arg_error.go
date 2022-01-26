package args

import "fmt"

const defaultMessage = "error with argument"

type ArgError struct {
	Err     error
	Context string
}

func (ae *ArgError) Unwrap() error { return ae.Err }

func (ae *ArgError) ErrorMessage() string {
	if len(ae.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, ae.Context, ae.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, ae.Err)
	}
}
