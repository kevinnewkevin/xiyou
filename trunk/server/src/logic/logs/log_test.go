package logs

import "testing"

func Test_IniString(t *testing.T) {
	Init()
	Error("%d%d%d%d%d%d", 1, 2, 3, 4, 5, 6)
	Error("%s", "This is error test 2")
	// Error("This is error test 3")
	// Error("This is error test 4")
	// Error("This is error test 5")
	// Error("This is error test 6")
	// Error("This is error test 7")

	Fini()
}
