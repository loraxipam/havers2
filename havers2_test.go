package havers2_test

import (
	"math"
	"testing"

	"github.com/loraxipam/havers2"
	"github.com/golang/geo/s2"
)

// Great circle calculations are really only valid for about four sig figs
// for Earth, anyway. Up this, (or should that be, "decrease this"?)
// for real spheres.
const ppm = 0.0001

var tests = []struct {
	p     havers2.Coord
	q     havers2.Coord
	outMi float64
	outKm float64
	outRa float64
	town  string
}{
	{
		havers2.Coord{Lat: -22.55, Lon: -43.12}, // Rio de Janeiro, Brazil
		havers2.Coord{Lat: 13.45, Lon: 100.28},  // Bangkok, Thailand
		9958,
		16026,
		2.515463,
		"Rio to Bangkok",
	},
	{
		havers2.Coord{Lat: -20.10, Lon: 57.30}, // Port Louis, Mauritius
		havers2.Coord{Lat: 0.57, Lon: 100.21},  // Padang, Indonesia
		3234,
		5205,
		0.817067,
		"Mauritius to Indonesia",
	},
	{
		havers2.Coord{Lat: 51.45, Lon: -1.15}, // Oxford, United Kingdom
		havers2.Coord{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		933,
		1501,
		0.235736,
		"Oxford to the Vatican",
	},
	{
		// These are antipodes
		havers2.Coord{Lat: 32.30, Lon: -64.77},  // Bermuda
		havers2.Coord{Lat: -32.30, Lon: 115.23}, // Perth, Australia
		12436,   // earthRadiusMi * pi
		20015,   // earthRadiusKm * pi
		3.14159, // pi
		"Bermuda to Perth",
	},
	{
		havers2.Coord{Lat: -22.34, Lon: 17.05}, // Windhoek, Namibia
		havers2.Coord{Lat: 51.56, Lon: 4.29},   // Rotterdam, Netherlands
		5164,
		8311,
		1.304548,
		"Namibia to the Netherlands",
	},
	{
		havers2.Coord{Lat: -63.24, Lon: -56.59}, // Esperanza, Argentina
		havers2.Coord{Lat: -8.50, Lon: 13.14},   // Luanda, Angola
		5069,
		8157,
		1.280482,
		"Argentina to Angola",
	},
	{
		havers2.Coord{Lat: 90.00, Lon: 0.00}, // North Pole
		havers2.Coord{Lat: 48.51, Lon: 2.21}, // Paris,  France
		2866,
		4613,
		0.724137,
		"Santa to Paris",
	},
	{
		havers2.Coord{Lat: -90.00, Lon: 0.00}, // South Pole
		havers2.Coord{Lat: 48.51, Lon: 2.21},  // Paris,  France
		9570,
		15401,
		2.417456,
		"Penguins to Paris",
	},
	{
		havers2.Coord{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		havers2.Coord{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		6262,
		10078,
		1.581873,
		"Turin to Malaysia",
	},
	{
		havers2.Coord{Lat: 45.71, Lon: -122.43}, // Hockinson, Washington, USA
		havers2.Coord{Lat: 29.13, Lon: -80.96},  // Wilbur-by-the-Sea, Florida, USA
		2510,
		4040,
		0.634270,
		"Washington to Florida",
	},
}

// TestHaversineDistance makes sure that the several great circle distance
// functions work properly
func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		// init the points
		input.p.Calc()
		input.q.Calc()
		mi, _ := math.Modf(havers2.DistanceMi(input.p, input.q))
		km, _ := math.Modf(havers2.DistanceKm(input.p, input.q))

		if input.outMi != mi || input.outKm != km {
			t.Errorf("fail: want %v %v -> %v %v got %g %g %s",
				input.p,
				input.q,
				input.outMi,
				input.outKm,
				mi,
				km,
				input.town,
			)
		}
	}
}

// TestIntAngle makes sure the internal angle calculation result is within tolerance
func TestIntAngle(t *testing.T) {
	for _, input := range tests {
		// init the points
		input.p.Calc()
		input.q.Calc()
		iAngle := havers2.IntAngle(input.p, input.q)
		delta := math.Abs(input.outRa-iAngle) / input.outRa

		if delta > ppm {
			t.Errorf("fail: want %v %v -> %v got %.6f, %.1f ppm %s",
				input.p,
				input.q,
				input.outRa,
				iAngle,
				delta*1000000,
				input.town,
			)
		}
	}
}

// TestPoleToPole makes sure the sum of internal angles of the north pole, a
// point and the south pole adds up to 180 degrees (pi radians)
func TestPoleToPole(t *testing.T) {
	var ll s2.LatLng
	var p  s2.Point
	reindeer := havers2.Coord{90, 0, ll, p}
	penguins := havers2.Coord{-90, 0, ll, p}
	reindeer.Calc()
	penguins.Calc()
	for _, input := range tests {
		// init the points
		input.p.Calc()
		input.q.Calc()
		nAngle := havers2.IntAngle(input.q, reindeer)
		sAngle := havers2.IntAngle(input.q, penguins)
		tAngle := nAngle + sAngle
		delta := math.Abs(tAngle-math.Pi) / math.Pi

		if delta > ppm {
			t.Errorf("fail: want %v -> %-.5f + %-.5f = Pi got %-.5f, %.1f ppm %s",
				input.q,
				nAngle,
				sAngle,
				tAngle,
				delta*1000000,
				input.town,
			)
		}
	}
}
