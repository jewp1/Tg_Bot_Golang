package e

import "fmt"

func Wrap(msg string, e error) error {
	return fmt.Errorf("%s: %w", msg, e)
}

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}
