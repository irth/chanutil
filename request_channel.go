package chanutil

import (
	"context"
)

type RequestChannel[Args any, Ret any] chan *Request[Args, Ret]

func (ch RequestChannel[Args, Ret]) Call(ctx context.Context, args Args) (*Ret, error) {
	req := &Request[Args, Ret]{
		ctx:  ctx,
		args: args,

		// channel queue size of 1 to make sure we don't stop execution
		// of the main loop
		ret: make(chan Result[Ret], 1),
	}

	// put and get fail if context gets cancelled
	err := Put(ctx, ch, req)
	if err != nil {
		return nil, err
	}

	ret, _, err := Get(ctx, req.ret)
	if err != nil {
		return nil, err
	}

	return ret.Unpack()
}
