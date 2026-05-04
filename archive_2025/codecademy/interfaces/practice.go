package main

type CalculateRentalCost() int {

}

type Car struct { //Car struct for 
// storing car usage data

	hours int // Hours the car has been used
	miles int // Miles the car has been driven

}

type Vehicle interface { // Interface defining methods required for vehicles
	CalculateRentalCost() int
}

type Scooter struct { //Scooter struct for storing scooter usage data
	hours int // Hours the scooter has been used
}

func (s Scooter) CalculateRentalCost() int {
	const hourlyRate int = 5 // Hourly rate
	return s.hours
}