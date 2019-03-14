package client
import (
	"github.com/satori/go.uuid"
)

type Client struct {
	id        uuid.UUID
	username  string
	password  string
	email     string
	payMethod string
}
