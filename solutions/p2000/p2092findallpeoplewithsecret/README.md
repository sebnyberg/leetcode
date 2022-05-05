# P2092

There are a couple of accepted solutions to this problem.

During the contest, I did managed to just barely pass using a quite slow heap-based solution where each edge was sorted on time (~4500ms).

# Simple solution (~760ms/26MB)

This solution was written to be clear, not fast. There are some easy perf. improvements that can be done to reduce runtime significantly.

Approach:

* Sort meetings by time
* Create list of meetings partitioned by time
* Keep track of which people know the secret
* For each unique time with meetings, iterate over meetings, adding new shared secrets until a loop did not result in any new people knowing about the secret.

```go
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// Partition meetings by time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	timeMeetings := make([][][]int, 1)
	timeMeetings[0] = append(timeMeetings[0], meetings[0])
	var timeIdx int
	for i := 1; i < len(meetings); i++ {
		if meetings[i][2] != meetings[i-1][2] {
			timeIdx++
			timeMeetings = append(timeMeetings, [][]int{})
		}
		timeMeetings[timeIdx] = append(timeMeetings[timeIdx], meetings[i])
	}

	knows := make([]bool, n)
	knows[0] = true
	knows[firstPerson] = true
	res := []int{0, firstPerson}
	// Iterate over meetings partitioned by timestamp
	for t := 0; t < len(timeMeetings); t++ {
		// Keep going through meetings until no new secrets were shared
		newSecret := true
		for newSecret {
			newSecret = false
			for _, meeting := range timeMeetings[t] {
				if knows[meeting[0]] && !knows[meeting[1]] {
					newSecret = true
					res = append(res, meeting[1])
					knows[meeting[1]] = true
				} else if !knows[meeting[0]] && knows[meeting[1]] {
					newSecret = true
					res = append(res, meeting[0])
					knows[meeting[0]] = true
				}
			}
		}
	}
	return res
}
```

# DSU (~390ms / 28MB)

Same approach as in previous exercise, but instead of looping over time-specific meetings until no more people are informed of the secret, use a Disjoint-Set Union (DSU).

DSU is a way to create sets where all elements of a set point to shared root. By cleverly picking the lowest possible root when joining sets together, it is easy to check whether people are in the group that knows about the secret - the root of their set will be zero.


```go
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// Set up DSU
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			root := find(parent[a])
			parent[a] = root // path compression
		}
		return parent[a]
	}
	union := func(a, b int) {
		aRoot, bRoot := find(a), find(b)
		if bRoot < aRoot {
			aRoot, bRoot = bRoot, aRoot // ensure that root of secret group will be 0
		}
		if aRoot != bRoot {
			parent[bRoot] = aRoot
		}
	}
	union(0, firstPerson)

	// Sort meetings by time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// Partition by time
	timeMeetings := make([][][]int, 1)
	timeMeetings[0] = append(timeMeetings[0], meetings[0])
	var timeIdx int
	for i := 1; i < len(meetings); i++ {
		if meetings[i][2] != meetings[i-1][2] {
			timeIdx++
			timeMeetings = append(timeMeetings, [][]int{})
		}
		timeMeetings[timeIdx] = append(timeMeetings[timeIdx], meetings[i])
	}

	// For each set of meetings for a given timestamp
	for _, meetings := range timeMeetings {
		// Add to DSU
		for _, meeting := range meetings {
			union(meeting[0], meeting[1])
		}

		// Reset entries which do not belong to root group
		for _, meeting := range meetings {
			if find(meeting[0]) != 0 {
				parent[meeting[0]] = meeting[0]
			}
			if find(meeting[1]) != 0 {
				parent[meeting[1]] = meeting[1]
			}
		}
	}

	// Add all nodes which are in secret group in DSU to result
	res := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		if find(i) == 0 {
			res = append(res, i)
		}
	}

	return res
}
```
