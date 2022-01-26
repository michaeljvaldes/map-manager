package mapdeploy

import "fmt"

const defaultMessage = "error deploying site"

type DeployError struct {
	Err     error
	Context string
}

func (de *DeployError) Unwrap() error { return de.Err }

func (de *DeployError) ErrorMessage() string {
	if len(de.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, de.Context, de.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, de.Err)
	}
}
