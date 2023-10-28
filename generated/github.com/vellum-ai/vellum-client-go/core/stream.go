package core

import (
	"errors"
	"net/http"
)

// Stream represents a stream of messages sent from a server.
type Stream[T any] struct {
	resp *http.Response
}

func (s Stream[T]) Recv() (T, error) {
	var result T
	return result, errors.New("unimplemented")
}

func (s Stream[T]) Close() error {
	return s.resp.Body.Close()
}
