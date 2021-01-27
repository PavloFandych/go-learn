package main

import (
	"fmt"
	"go-learn/core"
	"time"
)

func main() {
	var w core.Worker = &core.CustomWorker{}
	input := w.GenerateIntegerSlice(3000)
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

func sumPar(input []int32, result chan int64) {
	var total int64 = 0
	for _, value := range input {
		if isOdd(value) {
			total += int64(value)
		}
	}
	result <- total
}

func sumSeq(input []int32) int64 {
	var total int64 = 0
	for _, value := range input {
		if isOdd(value) {
			total += int64(value)
		}
	}
	return total
}

func isOdd(input int32) bool {
	return input&1 == 1
}
