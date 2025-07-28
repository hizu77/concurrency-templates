package fan

import (
	"context"
	"sync"

	"github.com/hizu77/concurrency-templates/internal/infra/pipeline"
)

func Fan[T, R any](
	ctx context.Context,
	fanLimit int,
	step func(ctx context.Context, t T) R,
	input <-chan T,
) <-chan R {
	chns := make([]<-chan R, 0, fanLimit)

	for range fanLimit {
		chns = append(chns, pipeline.GenericPipeline(ctx, input, step))
	}

	return fanIn(ctx, chns...)
}

func fanIn[T any](ctx context.Context, chans ...<-chan T) <-chan T {
	output := make(chan T)
	wg := &sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return
					case output <- v:
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
