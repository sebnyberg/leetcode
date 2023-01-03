package p0690employeeimportance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	var res int
	m := make(map[int]*Employee)
	for _, e := range employees {
		m[e.Id] = e
	}
	curr := []int{id}
	next := []int{}

	for len(curr) > 0 {
		next := next[:0]
		for _, x := range curr {
			res += m[x].Importance
			for _, i := range m[x].Subordinates {
				next = append(next, i)
			}
		}
		curr, next = next, curr
	}
	return res
}
