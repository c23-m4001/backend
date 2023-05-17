package geoip

import (
	"capstone/util"
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

type geoIp struct {
	reader *geoip2.Reader
}

func NewGeoIp(geoIpFilePath string) GeoIp {
	if !util.IsFileExist(geoIpFilePath) {
		panic(fmt.Sprintf("%s (GeoIp File) not exist", geoIpFilePath))
	}

	db, err := geoip2.Open(geoIpFilePath)
	if err != nil {
		panic(err)
	}

	return &geoIp{
		reader: db,
	}
}

func (g geoIp) ParseIP(_ip string) (*Payload, error) {
	ip := net.ParseIP(_ip)
	if ip == nil {
		return nil, ErrInvalidIpFormat
	}

	record, err := g.reader.City(ip)
	if err != nil {
		return nil, err
	}
	payload := Payload{
		Country:     record.Country.Names["en"],
		CountryCode: record.Country.IsoCode,
		City:        record.City.Names["en"],
		Timezone:    record.Location.TimeZone,
		Latitude:    record.Location.Latitude,
		Longitude:   record.Location.Longitude,
	}

	if len(record.Subdivisions) > 0 {
		payload.Subdivision = record.Subdivisions[0].Names["en"]
	}

	return &payload, nil
}

func (g *geoIp) Close() error {
	return g.reader.Close()
}
