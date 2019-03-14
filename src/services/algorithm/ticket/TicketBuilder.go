package ticket

import "services/algorithm/train"

type ticketBuilder struct {
	tickets ticket
}

func generatePrice(carriage train.TypeOfCarriage, routeDistance int) float32 {
	var totalPrice float32 = 0

	if (carriage == train.COMPARTMENT) {
		totalPrice = float32(routeDistance / 3.0)
	}
	if (carriage == train.BUSINESS) {
		totalPrice = float32(routeDistance / 5.0)

	}

	if (carriage == train.ECONOM) {
		totalPrice = float32(routeDistance / 10.0)
	}

	return totalPrice

}
