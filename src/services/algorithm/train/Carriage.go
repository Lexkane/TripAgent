package train

type Carriage struct {
	name           string
	amountOfPlaces int
	carriageType   TypeOfCarriage
}

func NewCarriage(name string, amountOfPlaces int, carriageType TypeOfCarriage) *Carriage {
	return &Carriage{name: name, amountOfPlaces: amountOfPlaces, carriageType: carriageType}
}
