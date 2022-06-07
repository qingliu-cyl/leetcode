package main

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	for i := 0; i < len(numbers) -1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}

	return numbers[0]
}

func main() {

}
