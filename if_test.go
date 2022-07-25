package functional

import "testing"

func TestIf(t *testing.T) {
	var success bool = false
	f1 := func() error { success = true; return nil }
	f2 := func() error { t.Errorf("If() called the wrong function"); return nil }

	If(true, f1, f2)
	switch {
	case !success:
		t.Errorf("If() didn't call the correct function")
	}
	If(false, f2, f1)
	switch {
	case !success:
		t.Errorf("If() didn't call the correct function")
	}

	success = false
	If(true, f1)
	switch {
	case !success:
		t.Errorf("If() didn't call the correct function")
	}

	success = false
	Unless(false, f1)
	switch {
	case !success:
		t.Errorf("If() didn't call the correct function")
	}
}
