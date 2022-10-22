package main

import (
	"errors"
	"fmt"
)

// FooError records an error and what caused it
type FooError struct {
	Op       string
	FooThing string
	Err      error
}

func ErrorCreator() error {
	return &FooError{
		Op:       "was eating miso soup",
		FooThing: "just happend",
		Err:      errors.New("bad fortune cookie"),
	}

}

func (e *FooError) Error() string { return e.Op + " " + e.FooThing + ": " + e.Err.Error() }

func (e *FooError) Unwrap() error { return e.Err }

func (e *FooError) Thing() string { return e.FooThing }

// The key here is to use "errors.As" to check if the error is of type FooError
func main() {
	if err := ErrorCreator(); err != nil {
		var fooError *FooError
		if errors.As(err, &fooError) {
			fmt.Println("Explain failure thing...:", fooError.Thing())
		} else {
			fmt.Println(err)
		}
	}

}
