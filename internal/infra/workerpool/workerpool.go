package workerpool

import (
	"context"
	"sync"
)

func Start[T, R any](
	ctx context.Context,
	input <-chan T,
	transform func(t T) R,
	workers int,
) <-chan R {
	output := make(chan R)
	wg := &sync.WaitGroup{}

	for range workers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-input:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return
					case output <- transform(v):
					}
				}
			}
		}()
	}

	go func() {
		defer close(output)
		wg.Wait()
	}()

	return output
}
