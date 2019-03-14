package journey

import (
	"../train"
	"time"
)

type Journey struct {
	name   string
	route  Route
	date   time.Time
	train  train.Train
	places map[string][][]int
}
