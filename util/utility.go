package util

func IsOddBinary(input int32) bool {
	return input&1 == 1
}

func IsOddMod(input int32) bool {
	return input%2 == 1
}

func SumPar(input []int32, result chan int64) {
	var total int64 = 0
	for _, value := range input {
		if IsOddBinary(value) {
			total += int64(value)
		}
	}
	result <- total
}

func SumSeq(input []int32) int64 {
	var total int64 = 0
	for _, value := range input {
		if IsOddBinary(value) {
			total += int64(value)
		}
	}
	return total
}
