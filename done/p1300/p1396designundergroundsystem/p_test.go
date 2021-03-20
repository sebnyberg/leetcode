package p1396designundergroundsystem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUndergroundSystem(t *testing.T) {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	// return 14.00000. There was only one travel from "Paradise" (at time 8) to "Cambridge" (at time 22)
	require.InEpsilon(t, float64(14),
		undergroundSystem.GetAverageTime("Paradise", "Cambridge"),
		0.001,
	)
	// return 11.00000. There were two travels from "Leyton" to "Waterloo", a customer with id=45 from time=3 to time=15 and a customer with id=27 from time=10 to time=20. So the average time is ( (15-3) + (20-10) ) / 2 = 11.00000
	require.InEpsilon(t, float64(11), undergroundSystem.GetAverageTime("Leyton", "Waterloo"), 0.001)
	undergroundSystem.CheckIn(10, "Leyton", 24)
	require.InEpsilon(t, float64(11), undergroundSystem.GetAverageTime("Leyton", "Waterloo"), 0.001) // return 11.00000
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	require.InEpsilon(t, float64(12), undergroundSystem.GetAverageTime("Leyton", "Waterloo"), 0.001) // return 12.00000
}

func TestUndergroundSystem2(t *testing.T) {
	sut := Constructor()
	sut.CheckIn(10, "Leyton", 3)
	sut.CheckOut(10, "Paradise", 8)
	require.InEpsilon(t, 5.0, sut.GetAverageTime("Leyton", "Paradise"), 0.01)
	sut.CheckIn(5, "Leyton", 10)
	sut.CheckOut(5, "Paradise", 16)
	require.InEpsilon(t, 5.5, sut.GetAverageTime("Leyton", "Paradise"), 0.01)
}

type checkIn struct {
	station string
	time    int
}

type checkOut struct {
	sum   int
	count int
}

type UndergroundSystem struct {
	checkIns  map[int]checkIn
	checkOuts map[[2]string][2]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkIns:  make(map[int]checkIn),
		checkOuts: make(map[[2]string][2]int),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkIns[id] = checkIn{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	c := this.checkIns[id]
	k := [2]string{c.station, stationName}
	this.checkOuts[k] = [2]int{
		this.checkOuts[k][0] + (t - c.time),
		this.checkOuts[k][1] + 1,
	}
	delete(this.checkIns, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	k := [2]string{startStation, endStation}
	return float64(this.checkOuts[k][0]) / float64(this.checkOuts[k][1])
}
