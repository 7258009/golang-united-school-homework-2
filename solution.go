package square
import(
"math"
)

type YourTypeNameHere int
	
const(
	SidesCircle YourTypeNameHere = 0
	SidesTriangle YourTypeNameHere = 3
	SidesSquare YourTypeNameHere = 4
)

func CalcSquare(sideLen float64, sidesNum YourTypeNameHere) float64 {

	if (sidesNum == SidesCircle){
		return (sideLen*sideLen/(4*math.Pi))
	}

	if (sidesNum == SidesSquare){
		return (sideLen*sideLen/16)
	}

	if (sidesNum == SidesTriangle){
		return (sideLen*math.Sqrt(sideLen*sideLen-sideLen*sideLen/4)/2)
	}
	
	return 0;
}

