package p2383minimumhoursoftrainingtowinacompetition

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var energySum int
	for _, e := range energy {
		energySum += e
	}
	var hoursOfTraining int
	if energySum+1 > initialEnergy {
		hoursOfTraining = energySum + 1 - initialEnergy
	}
	currExperience := initialExperience
	var experienceDelta int
	for _, e := range experience {
		if currExperience <= e {
			experienceDelta += e + 1 - currExperience
			currExperience += experienceDelta
		}
		currExperience += e
	}
	hoursOfTraining += experienceDelta
	return hoursOfTraining
}
