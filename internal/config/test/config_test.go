package test

import (
	"github.com/wlcmtunknwndth/AvitoTask/internal/config"
	"testing"
)

func TestMustLoad(t *testing.T) {
	conf := config.MustLoad()
	if conf == nil {
		t.Error("Error loading config")
	}
	t.Logf("loaded config: %+v", *conf)
}
