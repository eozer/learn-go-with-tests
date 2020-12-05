package ctx

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {
	// nop.
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
		/*
			ctx := r.Context()

			data := make(chan string, 1)

			go func() {
				data <- store.Fetch()
			}()

			select {
			// if we receive data as a result of goroutine doing store.Fetch()
			case d := <-data:
				fmt.Fprint(w, d)
			// if done is from request for cancelling the request.
			case <-ctx.Done():
				store.Cancel()
			}
		*/
	}
}
