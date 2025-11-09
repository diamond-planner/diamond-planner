package bsm

import "fmt"

type URLAllowlistError struct {
	URL     string
	Message string
}

func (e *URLAllowlistError) Error() string {
	return fmt.Sprintf("%s - %s", e.URL, e.Message)
}
