package main

import (
	"errors"

	"github.com/trebent/zerologr"
)

func main() {
	logger := zerologr.New(&zerologr.Opts{Caller: true, Console: false, V: 10})
	logger.Info("Hello, world!")
	logger.Error(errors.New("that's an error"), "Hello, world!")
	logger.V(1).Info("Hello, verbose world!")

	logger.WithValues("foo", "bar", "count", 113131).Info("Hello, world!")
	logger.WithName("foo").Info("Hello, world!")
	logger.WithName("foo").WithName("bar").Info("Hello, world!")

	zerologr.Set(zerologr.New(&zerologr.Opts{Caller: true, Console: false, V: 10}))
	zerologr.Info("Hello, world!")
	zerologr.Error(errors.New("that's an error"), "Hello, world!")
	zerologr.V(1).Info("Hello, verbose world!")
	zerologr.V(10).Info("Hello, even more verbose world!")
	zerologr.V(11).Info("This isn't shown!")
}
