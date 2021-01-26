package main

import (
	"fmt"
	"go-learn/core"
	"time"
)

func main() {
	var w core.Worker = &core.CustomWorker{}
	input := w.GenerateIntegerSlice(5000000)
	channel := w.GenerateIntegerChannel(2)

	defer w.Close(channel)

	//parallel execution
	parStart := time.Now()
	firstPart := input[:len(input)/2]
	secondPart := input[len(input)/2:]
	go sumPar(firstPart, channel)
	go sumPar(secondPart, channel)
	parResult := <-channel + <-channel
	parDuration := time.Since(parStart)

	//sequential execution
	seqStart := time.Now()
	seqResult := sumSeq(input)
	seqDuration := time.Since(seqStart)

	fmt.Printf("Result par: %d, duration: %d\n", parResult, parDuration)
	fmt.Printf("Result seq: %d, duration: %d\n", seqResult, seqDuration)
}

func sumPar(input []int, result chan int64) {
	var total int64 = 0
	for _, value := range input {
		total += int64(value)
	}
	result <- total
}

func sumSeq(input []int) int64 {
	var total int64 = 0
	for _, value := range input {
		total += int64(value)
	}
	return total
}
