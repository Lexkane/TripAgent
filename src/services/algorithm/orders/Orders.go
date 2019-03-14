package orders

import (
	"../client"
	"../ticket"
)

type Orders struct {
	orders map[client.Client][] ticket.Ticket
}
