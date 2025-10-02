package p3100numberofwaterbottles

func maxBottlesDrunk(numBottles int, numExchange int) int {
	res := numBottles
	empty := numBottles
	for empty >= numExchange {
		res++
		empty = empty - numExchange + 1
		numExchange++
	}
	return res
}
