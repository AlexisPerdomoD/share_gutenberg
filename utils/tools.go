package utils

import "context"

func Find(
	c context.Context,
	ch chan int,
	index int,
	current interface{},
	target interface{},
) {
	defer close(ch)
	select {
	case <-c.Done():
		return
	default:
		if current == target {
			ch <- index
		} else {
			ch <- -1
		}
	}
}
