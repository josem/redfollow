package redirect

import (
	"net/http"
	"time"
)

type Step struct {
	Code int
	Url  string
}

var client = &http.Client{
	Timeout: time.Second * 5,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func isRedirectionCode(code int) bool {
	httpRedirectCodes := []int{301, 302, 303, 307, 308}

	for _, httpRedirectCode := range httpRedirectCodes {
		if code == httpRedirectCode {
			return true
		}
	}
	return false
}

func computeRedirects(steps *[]Step) []Step {
	lastIndex := len(*steps) - 1
	stepToAnalyze := (*steps)[lastIndex]

	resp, error := client.Get(stepToAnalyze.Url)

	if error != nil {
		panic("Error")
	}

	stepToAnalyze.Code = resp.StatusCode

	(*steps)[lastIndex] = stepToAnalyze

	if isRedirectionCode(resp.StatusCode) {
		locationURL, e := resp.Location()
		if e != nil {
			panic(e)
		}
		nextStep := Step{Url: locationURL.String()}
		*steps = append(*steps, nextStep)

		if error != nil {
			panic(error)
		}

		computeRedirects(steps)
	}

	return *steps
}

func GetRedirects(url string) []Step {
	originalURL := Step{Url: url}
	steps := []Step{originalURL}

	return computeRedirects(&steps)
}
