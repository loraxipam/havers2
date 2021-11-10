# havers2 - Wrapping up geo/s2

This is a rewrite of the existing haversine library but uses geo/s2 for most internals.

See [loraxipam/circumpolar](https://github.com/loraxipam/circumpolar) and [loraxipam/weatherstem-cli](https://github.com/loraxipam/weatherstem-cli) for examples of usage.

## General info

This is a simple, in fact, primitive, wrapper around the basic geo/s2 library. It contains functions and constants that are Earth-centric, so you don't really need too many other things to calculate distances if you know latitudes and longitudes for locations on our bonny planet.

  - the Coord struct is the main one you'll use. Fill it with a lat/lon then pull the Calc() lever to create the supporting S2 LatLng and Point objects.
  - EarthRadius{Mi,Km,NM} are nominally accurate constants for earthy surface calculations.
  - Distance{Mi,Km,NM} are functions that return distances between two Coords, in those units. Sorry, you have to manage your own unit descriptors.
  - IntAngle uses the S2 library to determine the internal angle between two Coords. It's pretty much the core of this calculator.
  - Distance just multiples the above internal angle by a radius.
