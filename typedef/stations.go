package typedef

type DeparturesResponse struct {
	Timestamp string
	Available bool
	Station   string
	Response  []Departure
}

type Departure struct {
	Route       string
	Direction   string
	ArrivalTime int
}
