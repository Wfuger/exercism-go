package erratum

const testVersion = 2

func retry(attempts int, callback func() (Resource, error)) (Resource, error) {
	for i := 0; i < attempts; i++ {
		rsc, err := callback()
		if err == nil {
			return rsc, nil
		}
	}
	return callback()
}

// Use this func opens a resource and handles errors
func Use(o ResourceOpener, input string) (err error) {
	rsc, err := o()
	if err != nil {
		switch err.(type) {
		case TransientError:
			rsc2, err2 := retry(5, o)
			if err2 != nil {
				return err2
			}
			if rsc2 != nil {
				defer rsc2.Close()
				rsc2.Frob(input)
				return nil
			}
		default:
			return err
		}
	}
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case FrobError:
				rsc.Defrob(t.defrobTag)
				defer rsc.Close()
				err = t.inner
			case error:
				defer rsc.Close()
				err = t
			default:
				err = nil
			}
		}
	}()
	rsc.Frob(input)
	rsc.Close()
	return err
}
