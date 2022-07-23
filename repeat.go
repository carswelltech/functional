package iterable

type (
	Repeatable       func()
	RepeatableUnless func() error
	RepeatableWhile  func() (bool, error)
)

func Repeat(fct Repeatable, iterations uint) {
	for i := 0; i <= int(iterations); i++ {
		fct()
	}
}
func RepeatUnless(fct RepeatableUnless, iterations uint) error {
	for i := 0; i <= int(iterations); i++ {
		err := fct()
		switch {
		case err != nil:
			return err
		}
	}
	return nil
}
func RepeatWhile(fct RepeatableWhile) error {
	for {
		ok, err := fct()
		switch {
		case err != nil:
			return err
		case !ok:
			return nil
		}
	}
}
