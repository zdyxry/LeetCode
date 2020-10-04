package main

type ParkingSystem struct {
	m map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		m: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.m[carType] <= 0 {
		return false
	}
	this.m[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
