// Example demonstrates use of the havers2 Distance function.
package havers2_test

import (
	"fmt"

	"github.com/loraxipam/havers2"
)

func ExampleDistanceNM() {
	austin := havers2.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := havers2.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: -122.1430,
	}
	austin.Calc()
	paloAlto.Calc()

	nm := havers2.DistanceNM(austin, paloAlto)

	fmt.Printf("%.1f miles is a long walk to Silicon Valley.\n", nm)

	// Output: 1286.1 miles is a long walk to Silicon Valley.
}

func ExampleDistance() {
	austin := havers2.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := havers2.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: -122.1430,
	}
	austin.Calc()
	paloAlto.Calc()

	km := havers2.Distance(austin, paloAlto, 58230)

	fmt.Printf("%.1f km is a long walk to Silicon Valley on Saturn.\n", km)

	// Output: 21770.2 km is a long walk to Silicon Valley on Saturn.
}

func ExampleIntAngle() {
	austin := havers2.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := havers2.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: -122.1430,
	}
	austin.Calc()
	paloAlto.Calc()

	radians := havers2.IntAngle(austin, paloAlto)

	fmt.Printf("%.4f radians is a long walk to Silicon Valley, perhaps.\n", radians)

	// Output: 0.3739 radians is a long walk to Silicon Valley, perhaps.
}

func ExampleNmToKm() {
	fmt.Printf("One degree is 60 NM which is %.1f kilometers\n", havers2.NmToKm(60))

	// Output: One degree is 60 NM which is 111.1 kilometers
}

func ExampleDegrees() {
	// a planety thing
	type planet struct {
		name string
		r    float64
	}

	// km radius
	var solarSystem = []planet{
		{"the Sun", 695700},
		{"the Moon", 1737},
		{"Mercury", 2440},
		{"Venus", 6051},
		{"Earth", 6371},
		{"Mars", 3390},
		{"Jupiter", 69910},
		{"Saturn", 58230},
		{"Uranus", 25360},
		{"Neptune", 24620},
		{"the unit sphere", 1},
		// Sorry, Pluto, you got kicked to the curb, dude.
		// Your sorry ass is even smaller than the moon, anyway, so,
		// "Meh" is what I say to you. Sorry, Clyde. I love you, man.
	}

	// a winter seaside resort town
	wilbur := havers2.Coord{Lat: 29.13, Lon: -80.97}
	wilbur.Calc()
	fmt.Println("Wilbur-by-the-Sea is here:", wilbur.Lat, wilbur.Lon)
	fmt.Println(" which is", wilbur.S2LatLng.String())

	// it's antipode, a not winter not seaside not resort town
	summerhills := havers2.Coord{Lat: 45.71, Lon: -122.43}
	summerhills.Calc()
	fmt.Println("Summer Hills, however, is here:", summerhills.S2LatLng.String())

	// regardless of sphere, what's the angle?
	fmt.Printf(" That's an angular distance of %.4f radians\n", havers2.IntAngle(wilbur, summerhills))
	fmt.Printf(" or %.2f degrees\n", 180/3.14159*havers2.IntAngle(wilbur, summerhills))
	fmt.Printf(" so from Wilbur to Summer Hills is %.0f miles.\n", havers2.DistanceMi(wilbur, summerhills))

	// if this were a solar system object
	for _, v := range solarSystem {
		fmt.Printf("On %s, that would be %.1f miles or %.1f days of driving\n", v.name, havers2.Distance(wilbur, summerhills, havers2.NmToMi(havers2.KmToNm(v.r))), 6*v.r/6371)
	}

	// Output:
	// Wilbur-by-the-Sea is here: 29.13 -80.97
	//  which is [29.1300000, -80.9700000]
	// Summer Hills, however, is here: [45.7100000, -122.4300000]
	//  That's an angular distance of 0.6342 radians
	//  or 36.33 degrees
	//  so from Wilbur to Summer Hills is 2510 miles.
	// On the Sun, that would be 274138.8 miles or 655.2 days of driving
	// On the Moon, that would be 684.5 miles or 1.6 days of driving
	// On Mercury, that would be 961.5 miles or 2.3 days of driving
	// On Venus, that would be 2384.4 miles or 5.7 days of driving
	// On Earth, that would be 2510.5 miles or 6.0 days of driving
	// On Mars, that would be 1335.8 miles or 3.2 days of driving
	// On Jupiter, that would be 27547.9 miles or 65.8 days of driving
	// On Saturn, that would be 22945.4 miles or 54.8 days of driving
	// On Uranus, that would be 9993.0 miles or 23.9 days of driving
	// On Neptune, that would be 9701.4 miles or 23.2 days of driving
	// On the unit sphere, that would be 0.4 miles or 0.0 days of driving


}
