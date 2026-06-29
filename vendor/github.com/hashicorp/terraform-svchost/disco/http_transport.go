// Copyright IBM Corp. 2017, 2025

package disco

import (
	"net/http"
)

// DefaultUserAgent is the default User-Agent header value used in requests.
const DefaultUserAgent = "terraform-svchost/1.0"

// userAgentRoundTripper is an http.RoundTripper that adds a User-Agent header
// to requests.
type userAgentRoundTripper struct {
	innerRt   http.RoundTripper
	userAgent string
}

// newUserAgentTransport creates a new userAgentRoundTripper with the given ua string
func newUserAgentTransport(userAgent string, innerRt http.RoundTripper) http.RoundTripper {
	return &userAgentRoundTripper{
		innerRt:   innerRt,
		userAgent: userAgent,
	}
}

// RoundTrip implements the http.RoundTripper interface for userAgentRoundTripper
func (rt *userAgentRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", rt.userAgent)
	}

	return rt.innerRt.RoundTrip(req)
}
