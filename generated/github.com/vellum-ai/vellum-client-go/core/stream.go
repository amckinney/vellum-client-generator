package core

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Stream represents a stream of newline-delimited messages sent from a server.
//
// Python: https://github.com/vellum-ai/vellum-client-python/blob/1e577a183596da801faa1645de5b9ffc59da43ab/src/vellum/client.py#L111C40-L111C50
// Node: https://github.com/vellum-ai/vellum-client-node/blob/b05dc4c48c8ceb5dd117243d3a0dc554866cbb90/src/core/streaming-fetcher/StreamingFetcher.ts#L62
type Stream[T any] struct {
	reader io.Reader
	closer io.Closer
}

// NewStream constructs a new Stream from the given *http.Response.
func NewStream[T any](response *http.Response) *Stream[T] {
	return &Stream[T]{
		reader: bufio.NewReader(response.Body),
		closer: response.Body,
	}
}

func (s Stream[T]) Recv() (T, error) {
	// TODO: Parse the newline delimited lines until we hit the point
	// at which we need to decode the value.
	var message T
	if err := json.Unmarshal(nil /* TODO: This should be a formatted reader */, &message); err != nil {
		return message, err
	}
	return message, errors.New("unimplemented")
}

func (s Stream[T]) Close() error {
	return s.closer.Close()
}
