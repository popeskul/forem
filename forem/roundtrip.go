package forem

import (
	"fmt"
	"io"
	"net/http"
)

type LoggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l *LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[Request: %s %s]\n", r.Method, r.URL)
	return l.next.RoundTrip(r)
}
