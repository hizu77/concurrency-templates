package generator

import model "github.com/hizu77/concurrency-templates/internal/model/order"

func Generate(count int) <-chan model.Order {
	out := make(chan model.Order)

	go func() {
		defer close(out)

		for i := 1; i <= count; i++ {
			out <- model.Order{
				ID:    i,
				State: model.StateCreated,
			}
		}
	}()

	return out
}
