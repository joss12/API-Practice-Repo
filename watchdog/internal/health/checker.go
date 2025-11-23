package health

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

type CheckResult struct {
	Name    string        `json:"name"`
	OK      bool          `json:"ok"`
	Latency time.Duration `json:"latency_ms"`
	Error   string        `json:"error,omitempty"`
}

type Checker interface {
	Name() string
	Check(ctx context.Context) CheckResult
}

type HTTPChecker struct {
	name    string
	client  *http.Client
	method  string
	url     string
	timeout time.Duration
}

func NewHTTPChecker(name, method, url string, timeout time.Duration) *HTTPChecker {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).DialContext,
		TLSHandshakeTimeout:   timeout,
		ResponseHeaderTimeout: timeout,
		ForceAttemptHTTP2:     true,
	}
	return &HTTPChecker{
		name: name,
		client: &http.Client{
			Transport: tr,
			Timeout:   timeout * 2,
		},
		method:  method,
		url:     url,
		timeout: timeout,
	}
}

func (h *HTTPChecker) Name() string { return h.name }

func (h *HTTPChecker) Check(ctx context.Context) CheckResult {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, h.method, h.url, nil)
	resp, err := h.client.Do(req)
	dur := time.Since(start)

	res := CheckResult{Name: h.name, Latency: dur / time.Millisecond}
	if err != nil {
		res.OK = false
		res.Error = err.Error()
		return res
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		res.OK = true
		return res
	}
	res.OK = false
	res.Error = "non-2xx status: " + resp.Status
	return res
}

var ErrNoChecks = errors.New("no health checks configured")
