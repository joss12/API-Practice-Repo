package database

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	ConnectDB()
	if DB == nil {
		t.Fatal("database instance is nil")
	}
}
