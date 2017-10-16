package std

import "testing"

func Example() {

}

func TestLogger(t *testing.T){

	LogDebug("this is log %s", "test")
	LogWarning("this is log %s", "test")
	LogError("this is log %s", "test")
	LogFatal("this is log %s", "test")
	LogInfo("this is log %s", "test")

	LogBackup()
	LogDebug("this is log2 %s", "test")
	LogWarning("this is l2og %s", "test")
	LogError("this is l2og %s", "test")
	LogFatal("this is l2og %s", "test")
	LogInfo("this is lo2g %s", "test")

}