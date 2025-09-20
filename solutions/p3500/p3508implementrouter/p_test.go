package p3508implementrouter

import (
	"sort"
)

type packet struct {
	source, dest, timestamp int
}

type Router struct {
	packets []packet
	m       map[packet]struct{}
	destTs  map[int][]int
	sz      int
}

func Constructor(memoryLimit int) Router {
	return Router{
		packets: make([]packet, 0, memoryLimit),
		sz:      memoryLimit,
		m:       make(map[packet]struct{}, memoryLimit),
		destTs:  make(map[int][]int, memoryLimit),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	p := packet{
		source:    source,
		dest:      destination,
		timestamp: timestamp,
	}
	if _, exists := this.m[p]; exists {
		return false
	}
	if this.sz == len(this.packets) {
		this.ForwardPacket()
	}
	this.m[p] = struct{}{}
	this.destTs[p.dest] = append(this.destTs[p.dest], p.timestamp)
	this.packets = append(this.packets, p)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.packets) == 0 {
		return []int{}
	}
	p := this.packets[0]
	this.packets = this.packets[1:]
	this.destTs[p.dest] = this.destTs[p.dest][1:]
	delete(this.m, p)
	return []int{p.source, p.dest, p.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	start := sort.SearchInts(this.destTs[destination], startTime)
	end := sort.SearchInts(this.destTs[destination], endTime+1)
	return end - start
}
