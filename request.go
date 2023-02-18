package chanutil

import (
	"context"
	"fmt"
	"sync"
)

type Request[Args any, Ret any] struct {
	args Args

	ret  chan Result[Ret]
	ctx  context.Context
	once sync.Once
}

type Result[T any] struct {
	ok  *T
	err error
}

func (r Result[T]) Unpack() (*T, error) {
	return r.ok, r.err
}

func (r *Request[Args, Ret]) Context() context.Context {
	return r.ctx
}

func (r *Request[Args, Ret]) Args() Args {
	return r.args
}

var ErrResultSentTwice = fmt.Errorf("tried to send result twice")

func (r *Request[Args, Ret]) sendResult(res Result[Ret]) error {
	done := false
	r.once.Do(func() {
		r.ret <- res
		done = true
	})
	if !done {
		return ErrResultSentTwice
	}
	return nil
}

func (r *Request[Args, Ret]) Ok(ret Ret) error {
	return r.sendResult(Result[Ret]{ok: &ret})
}

func (r *Request[Args, Ret]) Err(err error) error {
	return r.sendResult(Result[Ret]{err: err})
}
