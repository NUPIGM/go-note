package notes2_test

import (
	"main/notes2"
	"testing"
)

func TestSync(t *testing.T) {
	notes2.Sync()
}
func TestWaiting(t *testing.T) {
	notes2.Waitting()
}

func TestGoroutine(t *testing.T) {
	notes2.Cond()
}

func TestCount(t *testing.T) {
	notes2.CountMutex()
}

func TestOnce(t *testing.T) {
	notes2.Once()
}
func TestMap(t *testing.T) {
	notes2.Map()
}
