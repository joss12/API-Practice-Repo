package metrics

import "testing"

func TestRegisterMetrics(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("metrics registration panicked: %v", r)
		}
	}()
	RegisterMetrics()
}
