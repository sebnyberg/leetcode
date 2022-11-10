package p0933numberofrecentcalls

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	for len(this.pings) > 0 && this.pings[0] < t-3000 {
		this.pings = this.pings[1:]
	}
	return len(this.pings)
}
