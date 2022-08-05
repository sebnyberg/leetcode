package p2358maximumnumberofgroupsneteringacompetition

func maximumGroups(grades []int) int {
	n := len(grades)
	var j int
	for i := 1; ; i++ {
		j += i
		if j > n {
			return i - 1
		}
	}
}
