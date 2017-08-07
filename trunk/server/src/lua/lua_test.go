package lua

import "testing"

func TestLlex(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)

		}
	}()
	t.Error(LUA_VSERSION)

	L := Open()
	RegistSystemAPI(L)
	r := LoadFile(L, "test.lua")
	r = CallP(L, 0, 1, 0)
	n := ToString(L, -1)
	t.Error(r, n)
}
