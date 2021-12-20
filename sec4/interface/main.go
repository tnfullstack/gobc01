// Interface
package main

import "fmt"

type Vihicle interface {
	TowReady() string
	OffRoadReady() string
	TrackRacing() string
	VihicleInfo()
}

type Car struct {
	Make     string
	Model    string
	Year     string
	TopSpeed int
	Track    bool
	OffRoad  bool
	Sound    string
	Electric bool
	TowPack  bool
	TowCap   int
}

type Truck struct {
	Make     string
	Model    string
	Year     string
	TopSpeed int
	Track    bool
	OffRoad  bool
	Sound    string
	Electric bool
	TowPack  bool
	TowCap   int
	TruckBed float64
}

func main() {

	// Toyota Tundra
	tundraTRD := Truck{
		Make:     "Toyota",
		Model:    "Tundra TRD Pro",
		Year:     "2022",
		TopSpeed: 180,
		Track:    false,
		OffRoad:  true,
		Sound:    "Loud",
		Electric: false,
		TruckBed: 6.5,
		TowPack:  true,
		TowCap:   12000,
	}

	// Lexus 350 RX
	lexus450RX := Car{
		Make:     "Toyota",
		Model:    "Lexus 450 RX",
		Year:     "2022",
		TopSpeed: 189,
		Track:    false,
		OffRoad:  false,
		Sound:    "quiet",
		Electric: true,
		TowPack:  false,
		TowCap:   3000,
	}

	PrintInfo(&lexus450RX)

	PrintInfo(&tundraTRD)

	t := tundraTRD.TowReady()
	fmt.Printf("%s, tow capacity is %d pounds\n", t, tundraTRD.TowCap)

	c := lexus450RX.TowReady()
	fmt.Printf("%s, tow capacity is %d pounds\n", c, lexus450RX.TowCap)

}

func (t *Truck) TowReady() string {
	if t.TowPack {
		return "This truck comes with tow package"
	} else {
		return "Tow package is not available"
	}
}

func (t *Truck) OffRoadReady() string {
	if t.OffRoad {
		return "This is an off-road ready truck."
	} else {
		return "This truck is not equit with off-road package"
	}
}

func (t *Truck) TrackRacing() string {
	if t.Track {
		return "This truck is track racing ready."
	} else {
		return "This truck is not for track racining."
	}
}

func (c *Car) TowReady() string {
	if c.TowPack {
		return "This car comes with tow package"
	} else {
		return "Tow package is not available"
	}
}

func (c *Car) OffRoadReady() string {
	if c.OffRoad {
		return "This is an off-road ready car."
	} else {
		return "This car is not equit with off-road package"
	}
}

func (c *Car) TrackRacing() string {
	if c.Track {
		return "This car is track-racing ready."
	} else {
		return "This car is not for track racining."
	}
}

func (t *Truck) VihicleInfo() {
	fmt.Println("VIHICLE INFORMATION")
	fmt.Println("-------------------")
	fmt.Printf("Make: %s\n", t.Make)
	fmt.Printf("Model: %s\n", t.Model)
	fmt.Printf("Year: %s\n", t.Year)
	fmt.Printf("Top Speed: %d\n", t.TopSpeed)
	fmt.Printf("Track Ready: %t\n", t.Track)
	fmt.Printf("Off Road: %t\n", t.OffRoad)
	fmt.Printf("Sound: %s\n", t.Sound)
	fmt.Printf("Electric: %t\n", t.Electric)
	fmt.Printf("TowPack: %t\n", t.TowPack)
	fmt.Printf("TowCap: %d\n", t.TowCap)
	fmt.Printf("Truck Bed: %.2f\n", t.TruckBed)
	fmt.Println("")
}

func (c *Car) VihicleInfo() {
	fmt.Println("VIHICLE INFORMATION")
	fmt.Println("-------------------")
	fmt.Printf("Make: %s\n", c.Make)
	fmt.Printf("Model: %s\n", c.Model)
	fmt.Printf("Year: %s\n", c.Year)
	fmt.Printf("Top Speed: %d\n", c.TopSpeed)
	fmt.Printf("Track Ready: %t\n", c.Track)
	fmt.Printf("Off Road: %t\n", c.OffRoad)
	fmt.Printf("Sound: %s\n", c.Sound)
	fmt.Printf("Electric: %t\n", c.Electric)
	fmt.Printf("TowPack: %t\n", c.TowPack)
	fmt.Printf("TowCap: %d\n", c.TowCap)
	// fmt.Printf("Truck Bed: %.2f\n", c.TruckBed)
	fmt.Println("")
}

func PrintInfo(v Vihicle) {
	v.VihicleInfo()
}
