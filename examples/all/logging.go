// nolint
package main

import (
	"errors"

	"github.com/trebent/zerologr"
)

func main() {
	logger := zerologr.New(&zerologr.Opts{Caller: true, Console: true, V: 10}).WithName("local")
	logger.Info("Hello, world!")
	logger.Error(errors.New("that's an error"), "Hello, world!")
	logger.V(1).Info("Hello, verbose world!")

	logger.WithValues("foo", "bar", "count", 113131).Info("Hello, world!")
	logger.WithName("foo").Info("Hello, world!")
	logger.WithName("foo").WithName("bar").Info("Hello, world!")

	zerologr.Set(zerologr.New(&zerologr.Opts{Caller: true, Console: true, V: 10}).WithName("pkg"))
	zerologr.Info("Hello, world!")
	zerologr.Error(errors.New("that's an error"), "Hello, world!")
	zerologr.V(1).Info("Hello, verbose world!")
	zerologr.V(10).Info("Hello, even more verbose world!")
	zerologr.V(11).Info("This isn't shown!")

	localPkg := zerologr.WithName("local")
	localPkg.Info("Hello, world!")
	localPkg = localPkg.WithValues("foo", "bar")
	localPkg.Info("Hello, world!")
	localPkg = zerologr.WithValues("count", 123)
	localPkg.Info("Hello, world!")

	zerologr.Set(zerologr.New(&zerologr.Opts{Caller: true, Console: true, V: 10}).WithName("depth"))
	zerologr.Info("Hello, world!")
	zerologr.Error(errors.New("that's an error"), "Hello, world!")
	zerologr.V(1).Info("Hello, verbose world!")
	zerologr.V(10).Info("Hello, even more verbose world!")
	zerologr.V(11).Info("This isn't shown!")

	zerologr.Set(zerologr.New(&zerologr.Opts{Caller: true, V: 10}).WithName("json"))
	zerologr.Info("Hello, JSON world!")
	zerologr.Error(errors.New("that's an error"), "Hello, world!")
}
