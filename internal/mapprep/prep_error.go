package mapprep

import "fmt"

const defaultMessage = "error creating site for map"

type PrepError struct {
	Err     error
	Context string
}

func (pe *PrepError) Unwrap() error { return pe.Err }

func (pe *PrepError) ErrorMessage() string {
	if len(pe.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, pe.Context, pe.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, pe.Err)
	}
}
