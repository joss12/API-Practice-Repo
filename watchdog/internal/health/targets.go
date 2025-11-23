package health

import (
	"context"
	"os"
	"strings"
	"time"
)

type Summary struct {
	Overall bool          `json:"overall"`
	Checks  []CheckResult `json:"checks"`
}

type Runner struct {
	checks []Checker
}

func NewRunnerFromEnv() *Runner {
	t := os.Getenv("WATCHDOG_TARGETS")
	if strings.TrimSpace(t) == "" {
		return &Runner{checks: nil}
	}
	timeout := 800 * time.Millisecond
	if v := strings.TrimSpace(os.Getenv("WATCHDOG_TIMEOUT_MS")); v != "" {
		if ms, err := time.ParseDuration(v + "ms"); err == nil && ms > 0 {
			timeout = ms
		}
	}
	var checks []Checker
	parts := strings.Split(t, ",")
	for _, p := range parts {
		fields := strings.Split(p, "|")
		if len(fields) != 3 {
			continue
		}
		name, method, url := fields[0], fields[1], fields[2]
		checks = append(checks, NewHTTPChecker(name, method, url, timeout))
	}
	return &Runner{checks: checks}
}

func (r *Runner) HasChecks() bool { return len(r.checks) > 0 }

func (r *Runner) Run(ctx context.Context) Summary {
	if len(r.checks) == 0 {
		return Summary{Overall: true, Checks: []CheckResult{}}
	}
	sum := Summary{Overall: true}
	for _, c := range r.checks {
		cr := c.Check(ctx)
		if !cr.OK {
			sum.Overall = false
		}
		sum.Checks = append(sum.Checks, cr)
	}
	return sum
}
