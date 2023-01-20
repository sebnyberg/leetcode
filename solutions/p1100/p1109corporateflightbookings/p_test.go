package p2

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	delta := make([]int, n+1)
	for _, b := range bookings {
		delta[b[0]-1] += b[2]
		delta[b[1]] -= b[2]
	}
	var v int
	for i := range res {
		v += delta[i]
		res[i] = v
	}
	return res
}
