package core

import "math/rand"

type CustomWorker struct {
}

func (c *CustomWorker) Close(channel chan int64) {
	close(channel)
}

func (c *CustomWorker) GenerateIntegerChannel(capacity int64) chan int64 {
	return make(chan int64, capacity)
}

func (c *CustomWorker) GenerateIntegerSlice(size int) []int {
	result := make([]int, 0, size)
	for index := 0; index < size; index++ {
		result = append(result, rand.Int())
	}
	return result
}
