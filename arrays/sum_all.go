package arrays

func SumAll(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return
}

func SumAllAtTail(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)

		} else {
			tails := number[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return
}
