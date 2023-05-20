package p1282groupthepeoplegiventhegroupsizetheybelongto

func groupThePeople(groupSizes []int) [][]int {
	groupPeople := make(map[int][]int)
	for i, g := range groupSizes {
		groupPeople[g] = append(groupPeople[g], i)
	}

	var res [][]int
	for g, p := range groupPeople {
		for i := 0; i < len(p); i += g {
			res = append(res, p[i:i+g])
		}
	}
	return res
}
