package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(slicesToSum ...[]int) []int {

	var sums []int

	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {

	var sums []int

	for _, slice := range slicesToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
