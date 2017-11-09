package logs

import "testing"

func TestLogger(t *testing.T) {

	Debug("this is log %s", "test")
	Warning("this is log %s", "test")
	Error("this is log %s", "test")
	Fatal("this is log %s", "test")
	Info("this is log %s", "test")

	Backup()
	Debug("this is log2 %s", "test")
	Warning("this is l2og %s", "test")
	Error("this is l2og %s", "test")
	Fatal("this is l2og %s", "test")
	Info("this is lo2g %s", "test")

}
