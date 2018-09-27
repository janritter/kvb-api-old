package typedef

type DeparturesResponse struct {
	Available bool
	Station   string
	Response  []Departure
}

type Departure struct {
	Route       string
	Direction   string
	ArrivalTime int
}
