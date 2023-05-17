package geoip

type Payload struct {
	Country     string
	CountryCode string
	Subdivision string
	City        string
	Timezone    string
	Latitude    float64
	Longitude   float64
}
