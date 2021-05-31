package test

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func a() error {
	return errors.Wrap(errors.New("root error"), "this is a")
}

func b() error {
	err := a()
	return errors.Wrap(err, "this is b")
}

func c() error {
	err := b()
	return errors.WithMessage(err, "this is c")
}

func TestError(t *testing.T) {
	err := c()
	if err != nil {
		fmt.Printf("Origin: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("Stack trace: \n%+v\n", err)
	}
}
