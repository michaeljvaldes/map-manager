package arguments

import "time"

type arguments struct {
	unminedPath string
	worldPath   string
	siteId      string
	deployToken string
	period      time.Duration
	startTime   time.Time
}

func (a arguments) Valid() (bool, []error) {
	errors := []error{}
	valid, err := a.validWorldPath()
	if !valid {
		errors = append(errors, err)
	}
	valid, err = a.validSiteId()
	if !valid {
		errors = append(errors, err)
	}
	valid, err = a.validDeployToken()
	if !valid {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return true, nil
	} else {
		return false, errors
	}
}

func (a arguments) validWorldPath() (bool, error) {
	return true, nil
}

func (a arguments) validSiteId() (bool, error) {
	return true, nil
}

func (a arguments) validDeployToken() (bool, error) {
	return true, nil
}
