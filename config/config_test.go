package config

import (
	"testing"
	"time"
)

func TestGetPollingInterval(t *testing.T) {
	got := GetPollingInterval()
	const want = 20 * time.Second

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

// TODO
func TestAskConfig(t *testing.T)     {}
func TestGetConfig(t *testing.T)     {}
func TestAskPass(t *testing.T)       {}
func TestStoreConfig(t *testing.T)   {}
func TestDestroyConfig(t *testing.T) {}
