package db_test

import (
	"main/demo/db"
	"testing"
)

func TestDb(t *testing.T) {
	db.LeveldbBasic()
}
