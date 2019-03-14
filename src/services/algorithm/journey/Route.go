package journey

type Route struct {
	begin string
	end   string
	stops map[string]int
}

func NewRoute(begin string, end string, stops map[string]int) *Route {
	return &Route{begin: begin, end: end, stops: stops}
}
