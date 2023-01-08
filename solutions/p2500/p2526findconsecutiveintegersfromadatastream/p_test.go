package p2526findconsecutiveintegersfromadatastream

type DataStream struct {
	val   int
	count int
	k     int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		val: value,
		k:   k,
	}
}

func (this *DataStream) Consec(num int) bool {
	if this.val != num {
		this.count = 0
		return false
	}
	this.count++
	return this.count >= this.k
}
