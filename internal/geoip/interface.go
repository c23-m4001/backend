package geoip

type GeoIp interface {
	ParseIP(ip string) (*Payload, error)
	Close() error
}
