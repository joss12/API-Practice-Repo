package status

import (
	"runtime"
	"time"

	"github.com/watchdog/internal/version"
)

type Snapshot struct {
	App        string    `json:"app"`
	Version    string    `json:"version"`
	Commit     string    `json:"commit"`
	BuiltAt    string    `json:"built_at"`
	GoVersion  string    `json:"go_version"`
	GoRoutines int       `json:"goroutines"`
	StartTime  time.Time `json:"start_time"`
	UptimeSec  int64     `json:"uptime_sec"`
}

type Tracker struct {
	start time.Time
}

func NewTracker() *Tracker { return &Tracker{start: time.Now()} }

func (t *Tracker) Snapshot() Snapshot {
	return Snapshot{
		App:        version.AppName,
		Version:    version.Version,
		Commit:     version.Commit,
		BuiltAt:    version.BuildDate,
		GoVersion:  runtime.Version(),
		GoRoutines: runtime.NumGoroutine(),
		StartTime:  t.start,
		UptimeSec:  int64(time.Since(t.start).Seconds()),
	}
}
