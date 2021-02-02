package main

import (
	"fmt"
	"go-learn/core"
	"go-learn/util"
	"time"
)

func main() {
	var w core.Worker = &core.CustomWorker{}
	input := w.GenerateIntegerSlice(3000000)
	channel := w.GenerateIntegerChannel(2)

	defer w.Close(channel)

	//parallel execution
	parStart := time.Now()
	firstPart := input[:len(input)/2]
	secondPart := input[len(input)/2:]
	go util.SumPar(firstPart, channel)
	go util.SumPar(secondPart, channel)
	parResult := <-channel + <-channel
	parDuration := time.Since(parStart)

	//sequential execution
	seqStart := time.Now()
	seqResult := util.SumSeq(input)
	seqDuration := time.Since(seqStart)

	fmt.Printf("Result par: %d, duration: %d\n", parResult, parDuration)
	fmt.Printf("Result seq: %d, duration: %d\n", seqResult, seqDuration)
}
