// Function with receiver
package main

import "fmt"

type Car struct {
	Make   string
	Model  string
	Year   string
	Wheel  int
	Door   int
	Seat   int
	Track  bool
	Speed  int
	Zero60 []float64
}

func main() {

	// Create a mustang
	mustang := Car{
		Make:   "Ford",
		Model:  "Mustang GT 5.0",
		Year:   "2021",
		Wheel:  4,
		Door:   2,
		Seat:   4,
		Track:  true,
		Speed:  189,
		Zero60: []float64{3.2, 3.5, 3.8},
	}

	// Create a corvett
	corvette := Car{
		Make:   "GM",
		Model:  "Corvette C8",
		Year:   "2021",
		Wheel:  4,
		Door:   2,
		Seat:   2,
		Track:  false,
		Speed:  189,
		Zero60: []float64{3.2, 3.1, 3.6},
	}

	// Printing car informaiton
	mustang.CarInfo()
	corvette.CarInfo()

	// Calculate car average speed
	mRaceAverage := mustang.RaceAverageSpeed()
	cRaceAverage := corvette.RaceAverageSpeed()
	fmt.Printf("The Mustange GT 5.0 average (0 to 60) winning record is %.2f\n", mRaceAverage)
	fmt.Printf("The Covette C8 average (0 to 60) winning record is %.2f\n", cRaceAverage)
	fmt.Println("")

	// Compare Mustange and Covette winning recoard

	if mRaceAverage < cRaceAverage {
		fmt.Printf("The %s is the winner\n", mustang.Model)
	} else {
		fmt.Printf("The %s is the winner\n", corvette.Model)
	}
}

func (c *Car) CarInfo() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model: ", c.Model)
	fmt.Println("Year built: ", c.Year)
	fmt.Println("Whell: ", c.Wheel)
	fmt.Println("Door: ", c.Door)
	fmt.Println("Seat: ", c.Seat)
	fmt.Println("Track: ", c.Track)
	fmt.Println("Speed: ", c.Speed)
	fmt.Printf("The last %d (0 to 60) racing record are %.2f\n", len(c.Zero60), c.Zero60)
	fmt.Println("")
}

func (c *Car) RaceAverageSpeed() float64 {
	a := 0.0
	l := len(c.Zero60)
	for _, v := range c.Zero60 {
		a += v
	}
	return a / float64(l)
}
