package train

type Train struct {
	listofCarriages [] Carriage
}

func NewTrain(name string, amountOfPlaces int, carriageType TypeOfCarriage) *Train {
	carriage := NewCarriage(name, amountOfPlaces, carriageType)
	return &Train{listofCarriages: append(make([]Carriage, 0), *carriage)}
}
