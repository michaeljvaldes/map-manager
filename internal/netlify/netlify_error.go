package netlify

import "fmt"

const defaultMessage = "error sending request to netlify"

type NetlifyError struct {
	Err     error
	Context string
}

func (ne *NetlifyError) Unwrap() error { return ne.Err }

func (ne *NetlifyError) ErrorMessage() string {
	if len(ne.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, ne.Context, ne.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, ne.Err)
	}
}
