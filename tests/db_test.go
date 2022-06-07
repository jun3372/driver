package tests

import (
	"testing"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	dsn = "dm://SYSDBA:SYSDBA@localhost:5236"
)

func TestMakeDB(t *testing.T) {
	var err error
	if db, err = MakeDB(dsn); err != nil {
		t.Fatal(err)
	}
	t.Parallel()
}
