# chanutil

[![Go Reference](https://pkg.go.dev/badge/github.com/irth/chanutil.svg)](https://pkg.go.dev/github.com/irth/chanutil) [![Tests](https://github.com/irth/chanutil/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/irth/chanutil/actions/workflows/go.yml)

Some utilities for interacting with channels that I use in my code.

- `Get` - receive from a channel in a blocking manner, but return if a context expires.
- `Put` - the same, for sending to a channel.

Currently, it's so small, that it probably makes no sense to use it. Feel free
to just steal the code you want, no need to introduce unnecessary dependency
on this lib. It's MIT.
