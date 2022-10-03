package p2424longestuploadedprefix

type LUPrefix struct {
	n      int
	prefix int
	seen   []bool
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		n:      n,
		seen:   make([]bool, n+1),
		prefix: 1,
	}
}

func (this *LUPrefix) Upload(video int) {
	if this.seen[video] == true {
		return
	}
	this.seen[video] = true
	for this.prefix <= this.n && this.seen[this.prefix] {
		this.prefix++
	}
}

func (this *LUPrefix) Longest() int {
	return this.prefix - 1
}
