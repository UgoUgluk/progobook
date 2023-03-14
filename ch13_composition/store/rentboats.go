package store

//Crew  describes a boat crew
type Crew struct {
	Captain, FirstOfficer string
}

// RentalBoat expands boat
type RentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

// NewRentalBoat create new reantal boat
func NewRentalBoat(
	name string,
	price float64,
	capacity int,
	motorized,
	crewed bool,
	captain, firstOfficer string,
) *RentalBoat {
	return &RentalBoat{
		NewBoat(name, price, capacity, motorized),
		crewed,
		&Crew{captain, firstOfficer},
	}
}
