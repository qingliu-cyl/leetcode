package main

import "fmt"

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	hours := 0

	for i := 0; i < len(energy); i++ {
		if initialEnergy > energy[i] {
			initialEnergy -= energy[i]
		} else {
			hours += energy[i] - initialEnergy + 1
			initialEnergy = 1
		}

		if initialExperience <= experience[i] {
			hours += experience[i] - initialExperience + 1
			initialExperience += experience[i] - initialExperience + 1
		}

		initialExperience += experience[i]
	}
	return hours
}

func main()  {
	fmt.Println(minNumberOfHours(1,1, []int{1,1,1,1}, []int{1,1,1,50}))
}