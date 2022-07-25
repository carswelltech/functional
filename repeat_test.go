package functional

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	a := 0
	f := func() error {
		a++
		if a >= 100 {
			return fmt.Errorf("error")
		}
		return nil
	}
	err := Repeat(f, 50)
	e := 50
	if err != nil {
		t.Errorf("Repeat(): Unexpected Error %v", err)
	}
	if a != e {
		t.Errorf("Repeat(): expected %v, got %v", e, a)
	}

	a = 0
	err = Repeat(f, 101)
	e = 100
	if err == nil {
		t.Errorf("Repeat(): expected error, got none")
	}
	if a != e {
		t.Errorf("Repeat(): expected %v, got %v", e, a)
	}

}
func TestRepeatWhile(t *testing.T) {
	a := 0
	f := func() (bool, error) {
		a++
		switch {
		case a >= 100:
			return false, nil
		default:
			return true, nil
		}
	}
	f2 := func() (bool, error) {
		a++
		switch {
		case a >= 100:
			return true, fmt.Errorf("error")
		default:
			return true, nil
		}
	}

	e := 100
	err := RepeatWhile(f)
	if err != nil {
		t.Errorf("RepeatWhile(): Unexpected Error")
	}
	if a != e {
		t.Errorf("RepeatWhile(): expcted %v, got %v", e, a)
	}

	a = 0
	err = RepeatWhile(f2)
	if err == nil {
		t.Errorf("RepeatWhile(): expected error, got none")
	}
	if a != e {
		t.Errorf("RepeatWhile(): expected %v, got %v", e, a)
	}
}
