package args

import "time"

type Arguments struct {
	UnminedPath string
	WorldPath   string
	SiteId      string
	DeployToken string
	Period      time.Duration
	StartTime   time.Time
}

func (a Arguments) Valid() (bool, []error) {
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

func (a Arguments) validWorldPath() (bool, error) {
	return true, nil
}

func (a Arguments) validSiteId() (bool, error) {
	return true, nil
}

func (a Arguments) validDeployToken() (bool, error) {
	return true, nil
}
