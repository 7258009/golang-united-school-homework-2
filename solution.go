package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

//func CalcSquare(sideLen float64, sidesNum #yourTypeNameHere#) float64 {
//}


import(
"fmt"
"math"
)

//package square
// Define custom int type to hold sides number and update CalcSquare signature by replacing 
//yourTypeNameHere
// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

	type YourTypeNameHere int
	
	const(
		SidesCircle YourTypeNameHere = 0
		SidesTriangle YourTypeNameHere = 3
		SidesSquare YourTypeNameHere = 4
	)

func CalcSquare(sideLen float64, sidesNum YourTypeNameHere) float64 {

	if (sidesNum == SidesCircle){
		return sideLen*sideLen/(4*math.Pi)
	}

	if (sidesNum == SidesSquare){
		return (sideLen*sideLen/16)
	}

	if (sidesNum == SidesTriangle){
		return (sideLen*sideLen)/(3*3*2*2)
	}
	
	return 0;
}

