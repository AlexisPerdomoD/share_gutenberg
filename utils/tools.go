package utils

import "context"

func Find(
	ctx context.Context,
	ch chan int,
	collection []int,
	target int,
) {
	defer close(ch)

	for index, current := range collection {
		select {
		case <-ctx.Done():
			return
		default:
			if current == target {
				ch <- index
			} else {
				ch <- -1
			}
		}
	}
}
