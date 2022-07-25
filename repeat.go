package functional

type (
	Repeatable      func() error
	RepeatableWhile func() (bool, error)
)

/*
	Repeat accepts a function argument, and a uint argument.  It calls the function argument the number of times indicated by the uint argument, or until the function argument returns a non-nil error, whichever comes first.
*/
func Repeat(fct Repeatable, iterations uint) error {
	for i := 1; i <= int(iterations); i++ {
		err := fct()
		switch {
		case err != nil:
			return err
		}
	}
	return nil
}

/*
	RepeatWhile accepts a function argument, which it calls until the function either returns false, or a non-nil error
*/
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
