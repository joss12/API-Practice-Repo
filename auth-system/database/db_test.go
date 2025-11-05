package database

import (
	"os"
	"testing"

	"github.com/auth-system/config"
)

func TestConnectDB(t *testing.T) {
	os.Setenv("DB_URL", "root:password@tcp(localhost:3306)/authdb?parseTime=true")
	config.LoadConfig()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ConnectDB panicked: %v", r)
		}
	}()

	ConnectDB()

	if DB == nil {
		t.Fatalf("DB should not be nil after connection")
	}
}
