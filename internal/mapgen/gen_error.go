package mapgen

import "fmt"

const defaultMessage = "Error generating maps"

type GenError struct {
	Err     error
	Context string
}

func (ge *GenError) Unwrap() error { return ge.Err }

func (ge *GenError) ErrorMessage() string {
	if len(ge.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, ge.Context, ge.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, ge.Err)
	}
}
