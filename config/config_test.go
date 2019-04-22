package config

import (
	"flag"
	"testing"
	"time"
)

func TestGetPollingInterval(t *testing.T) {
	got := GetPollingInterval()
	const want = time.Duration(15 * time.Second)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestGetPollingInterval2(t *testing.T) {
	// TODO Figure out the configured case with the `-poll=n` flag
	flag.Parse()
}

// TODO
func TestAskConfig(t *testing.T)     {}
func TestGetConfig(t *testing.T)     {}
func TestAskPass(t *testing.T)       {}
func TestStoreConfig(t *testing.T)   {}
func TestDestroyConfig(t *testing.T) {}
