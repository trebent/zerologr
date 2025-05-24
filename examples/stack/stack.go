// nolint
package main

import "github.com/trebent/zerologr"

func main() {
	logger := zerologr.New(&zerologr.Opts{Caller: true, Console: true, V: 10})
	logger.Info("Hello, world!")

	zerologr.Set(logger.WithName("pkg"))

	zerologr.Info("Hello, world!")

	pkgLogger := zerologr.WithName("local")
	pkgLogger.Info("Hello, world!")
}
