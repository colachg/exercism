package elon

import (
	"strconv"
)

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// Track implements a race track.
type Track struct {
	distance int
}

// CreateCar creates a new car with given specifications.
func CreateCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery: 100,
		distance: 0,
	}
}

// CreateTrack creates a new track with given distance.
func CreateTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time.
func (car *Car) Drive() {
	car.battery -= car.batteryDrain
	car.distance = car.speed
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(track Track) bool {
	return track.distance / car.speed * car.batteryDrain <= car.battery
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	return "Driven " +strconv.Itoa(car.distance) + " meters"
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}
