package p1603designparkingsystem

type ParkingSystem struct {
	size [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		size: [3]int{big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.size[carType-1] == 0 {
		return false
	}
	this.size[carType-1]--
	return true
}
