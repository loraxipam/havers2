/*
Package havers2 provides the great circle distance between two points
on a sphere. Several functions return great circle distances of the Earth.
Points are identified by latitude and longitude in degrees, and distance
results are returned in nautical miles.  Functions are also provided to
return the distance in kilometers or statute miles, or the geocentric
angular separation between those points.

For a complete toolset of polar and cartesian calculations, see the geo/r? and
geo/s? libraries, upon which this library is now based.

The below example shows how to calculate the shortest path between two
coordinates on the surface of the Earth.

    package main

    import (
        "fmt"
        "math"

        "github.com/loraxipam/havers2"
    )

    func main() {
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
        fmt.Printf("it's only %.1f miles on the moon, though.\n", havers2.Distance(austin, paloAlto, 937.9))
        fmt.Printf("That's an angle of %.1f degrees\n", 180/math.Pi*havers2.IntAngle(austin, paloAlto))
    }
*/
package havers2
