package havers2

import (
	"github.com/golang/geo/s2"
)

const (
	// 1NM = 1852m exactly
	EarthRadiusMi = 3958.8 // radius of the earth in miles.
	EarthRadiusKm = 6371.0 // radius of the earth in kilometers.
	EarthRadiusNM = 3440.1 // radius of the earth in nautical miles
)

// Coord represents a lat/long geographic coordinate, usually in degrees +EN/-WS.
type Coord struct {
	Lat float64		// degrees 90..-90
	Lon float64		// degrees 180..-180
	S2LatLng s2.LatLng
	S2Point s2.Point
}

// Calc() populates the LatLng and Point structs in a Coord
func (c *Coord) Calc() {
	c.S2LatLng = s2.LatLngFromDegrees(c.Lat, c.Lon)
	c.S2Point = s2.PointFromLatLng(c.S2LatLng)
}

// NmToMi converts from nautical miles to statute miles.
func NmToMi(d float64) float64 {
	return d * EarthRadiusMi / EarthRadiusNM
}

// NmToKm converts from nautical miles to kilometers.
func NmToKm(d float64) float64 {
	return d * EarthRadiusKm / EarthRadiusNM
}

// KmToNm converts from kilometers to nautical miles.
func KmToNm(d float64) float64 {
	return d * EarthRadiusNM / EarthRadiusKm
}

// IntAgle calculates the internal angle between two coordinates on a surface
// and returns the result in radians
func IntAngle(p, q Coord) float64 {
	return p.S2Point.Distance(q.S2Point).Radians()
}

// Distance calculates the shortest (aka great circle) arc between two coordinates on a sphere
// of a given radius and returns the resulting great circle arc length
func Distance(p, q Coord, radius float64) (gc float64) {
	return IntAngle(p, q) * radius
}

// DistanceMi calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in statute miles.
func DistanceMi(p, q Coord) (mi float64) {
	return Distance(p, q, EarthRadiusMi)
}

// DistanceKm calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in kilometers.
func DistanceKm(p, q Coord) (km float64) {
	return Distance(p, q, EarthRadiusKm)
}

// DistanceNM calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in nautical miles.
func DistanceNM(p, q Coord) (nm float64) {
	return Distance(p, q, EarthRadiusNM)
}
