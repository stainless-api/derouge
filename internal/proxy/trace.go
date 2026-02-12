package proxy

import "time"

type RequestTrace struct {
	Start time.Time

	// Phase durations
	DenyListCheck  time.Duration
	JWEDecrypt     time.Duration
	PathParse      time.Duration
	HostMatch      time.Duration
	RequestBuild   time.Duration
	UpstreamRoundTrip time.Duration
	ResponseStream time.Duration

	// Total wall-clock time for the entire request
	Total time.Duration

	// Metadata
	Method         string
	TargetHost     string
	TargetPath     string
	StatusCode     int
	CredentialCount int
	ResponseBodyBytes int64
}

type TraceCallback func(RequestTrace)
