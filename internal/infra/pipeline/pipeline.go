package pipeline

import (
	"context"
)

func GenericPipeline[T, R any](
	ctx context.Context,
	input <-chan T,
	step func(ctx context.Context, t T) R,
) <-chan R {
	output := make(chan R)

	go func() {
		defer close(output)

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
				case output <- step(ctx, v):
				}

			}
		}
	}()

	return output
}
