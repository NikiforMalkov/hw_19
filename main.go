package main

import "fmt"

type segment struct {
	sum   int
	start int
	end   int
}

func (s *segment) String() string {
	return fmt.Sprintf("started at: %d, ended at: %d", s.start, s.end)
}

func main() {
	data := []int{10, -3, -12, 8, 42, 1, -7, 0, 3}
	result := findClosedSubarrayWithMaxSum(data)
	fmt.Println("Result ", result.String())
}

func findClosedSubarrayWithMaxSum(array []int) segment {
	max := segment{
		sum:   0,
		start: 0,
		end:   0,
	}

	currentSegment := segment{
		sum:   0,
		start: 0,
		end:   0,
	}

	//Чисто технически, если наша сумма уменьшилась, тогда мы нашли конец отрезка
	for key, value := range array {
		sum := currentSegment.sum + value
		if currentSegment.sum > sum {
			currentSegment = segment{
				sum:   0,
				start: key + 1,
				end:   0,
			}
		}
		currentSegment.sum = sum
		if currentSegment.sum > max.sum {
			currentSegment.end = key
			max = currentSegment
		}
	}
	return max
}

//findSubarrayWithMaxSum Первый пример
func findSubarrayWithMaxSum(array []int) segment {
	max := segment{
		sum:   0,
		start: 0,
		end:   0,
	}

	currentSegment := segment{
		sum:   0,
		start: 0,
		end:   0,
	}

	for key, value := range array {
		if value < 0 {
			if max.sum < currentSegment.sum {
				max = currentSegment
			}
			currentSegment = segment{
				sum:   0,
				start: 0,
				end:   0,
			}
			currentSegment.start = key + 1
		} else {
			currentSegment.sum += value
			fmt.Println(key - 1)
			currentSegment.end = key
		}
	}
	if currentSegment.sum > max.sum {
		currentSegment.end = len(array) - 1
		max = currentSegment
	}
	return max
}
