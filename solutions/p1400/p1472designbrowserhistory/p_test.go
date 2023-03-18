package p1472designbrowserhistory

type BrowserHistory struct {
	pages []string
	i     int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		pages: []string{homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.pages = this.pages[:this.i+1]
	this.pages = append(this.pages, url)
	this.i++
}

func (this *BrowserHistory) Back(steps int) string {
	this.i = max(0, this.i-steps)
	return this.pages[this.i]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.i = min(len(this.pages)-1, this.i+steps)
	return this.pages[this.i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
