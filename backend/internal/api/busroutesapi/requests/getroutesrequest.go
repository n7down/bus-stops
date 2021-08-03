package requests

import (
	"net/url"
	"regexp"
)

type GetRoutesRequest struct {
	Stop int    `json:"stop" binding:"required"`
	Time string `json:"time" binding:"required"`
}

func (r *GetRoutesRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("([0-9][0-2]):([0-9][0-9])(AM|PM)")
	isValidTime := regex.MatchString(r.Time)

	if r.Stop < 1 {
		errs.Add("stop", "The stop field should be a positive non-zero value!")
	}

	if r.Stop > 10 {
		errs.Add("stop", "The stop should be between 1 and 10")
	}

	if !isValidTime {
		errs.Add("time", "The time format is not valid!")
	}

	return errs
}
