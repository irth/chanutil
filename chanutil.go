package chanutil

import (
	"context"
)

// Put tries to put a message in a channel, blocking until it succeeds or the
// context gets cancelled or expires. Returns ctx.Err() if the message was not
// sent successfuly.
func Put[T any](ctx context.Context, ch chan<- T, v T) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case ch <- v:
		return nil
	}
}

// Get tries to read a message from a channel, blocking until it succeeds or the
// context gets cancelled or expires.
//
// Returns a pointer to the message if it was received successfuly (nil
// otherwise), a bool equal to true if the channel isn't closed, and ctx.Err()
// if the context expired/got cancelled.
func Get[T any](ctx context.Context, ch <-chan T) (*T, bool, error) {
	select {
	case <-ctx.Done():
		return nil, false, ctx.Err()
	case v, more := <-ch:
		return &v, more, nil
	}
}
