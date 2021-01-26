package core

type Worker interface {
	IntegerSliceGenerator
	IntegerChannelGenerator
	IntegerChannelCloser
}

type IntegerSliceGenerator interface {
	GenerateIntegerSlice(size int) []int
}

type IntegerChannelGenerator interface {
	GenerateIntegerChannel(capacity int64) chan int64
}

type IntegerChannelCloser interface {
	Close(channel chan int64)
}
